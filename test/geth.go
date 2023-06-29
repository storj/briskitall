package test

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"storj.io/briskitall/internal/eth"
)

type geth struct {
	NodeURL        string
	Client         eth.Client
	ChainID        *big.Int
	RootTransactor *bind.TransactOpts

	containerName string
}

type testLogOut struct {
	t *testing.T
}

func (o *testLogOut) Write(b []byte) (int, error) {
	o.t.Logf("%s", string(b))
	return len(b), nil
}

// fund uses the geth JS console to fund addresses with 10 ETH each
func (g *geth) Fund(t *testing.T, addresses ...common.Address) {
	out := &testLogOut{t}
	cmds := []string{
		"miner.stop",
	}

	for _, address := range addresses {
		cmds = append(cmds, fmt.Sprintf(`eth.sendTransaction({from: eth.accounts[0], to: %q, value: web3.toWei(10, "ether")})`, address))
	}

	cmds = append(cmds,
		"miner.start",
		"admin.sleepBlocks(1)",
	)

	scriptCmd := exec.Command("docker", "exec", g.containerName,
		"geth", "attach",
		"--datadir", ".",
		"--exec",
		strings.Join(cmds, ";"),
	)
	t.Logf("Fund: args=%q", scriptCmd.Args)
	scriptCmd.Stdout = out
	scriptCmd.Stderr = out
	err := scriptCmd.Run()
	require.NoError(t, err, "failed to fund %q", addresses)
}

// runGeth starts geth in a docker container with a random ID. The geth RPC port
// is bound to an ephemeral port on the host.
func runGeth(t *testing.T) *geth {
	containerName := os.Getenv("BRISKITALL_TEST_CONTAINER")
	if containerName == "" {
		containerName = fmt.Sprintf("briskitall-test-%d", rand.Uint64())
		t.Logf("Starting geth container %q", containerName)
		startContainer(t, containerName)
	} else {
		t.Logf("Using existing geth container %q", containerName)
	}

	// Figure out the ephemeral port.
	port := getLocalPort(t, containerName)
	t.Logf("Ephemeral local port for container %q: %s", containerName, port)
	url := "http://localhost:" + port

	client, err := eth.Dial(context.Background(), url)
	require.NoError(t, err)
	t.Cleanup(client.Close)

	// Wait for the geth node to report its chain ID, as a way of knowing that
	// the node has initialized.
	var chainID *big.Int
	require.Eventually(t, func() bool {
		var err error
		chainID, err = client.ChainID(context.Background())
		return err == nil
	}, time.Second*5, time.Millisecond*20, "timed out waiting for node to start")

	rootTransactor, err := bind.NewKeyedTransactorWithChainID(rootKey, chainID)
	require.NoError(t, err)

	return &geth{
		NodeURL:        url,
		Client:         client,
		ChainID:        chainID,
		RootTransactor: rootTransactor,

		containerName: containerName,
	}
}

func startContainer(t *testing.T, containerName string) {
	runCmd := exec.Command("docker", "run",
		// docker flags
		"-d", "--rm", "--name", containerName, "-p", ":8545", "ethereum/client-go",
		// geth flags
		"--dev", "--datadir", ".", "--http", "--http.addr", "0.0.0.0", "--http.api", "eth,web3,net,debug")
	err := runCmd.Run()
	require.NoError(t, err, "failed to start geth container")
	t.Cleanup(func() {
		require.NoError(t, exec.Command("docker", "stop", containerName, "-s", "kill").Run())
	})
}

func getLocalPort(t *testing.T, containerName string) string {
	stdout := new(bytes.Buffer)
	portCmd := exec.Command("docker", "port", containerName, "8545/tcp")
	portCmd.Stdout = stdout
	err := portCmd.Run()
	require.NoError(t, err, "failed to get geth container port")

	scanner := bufio.NewScanner(stdout)
	scanner.Scan()
	require.NoError(t, scanner.Err(), "failed to read geth container port")

	_, port, found := strings.Cut(strings.TrimSpace(scanner.Text()), ":")
	require.True(t, found, "found port")
	return port
}
