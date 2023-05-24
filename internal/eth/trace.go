package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func TraceTransaction(ctx context.Context, nodeURL string, hash common.Hash) error {
	params := struct {
		ID      int64
		Version string `json:"jsonrpc"`
		Method  string
		Params  []interface{}
	}{
		ID:      1,
		Version: "2.0",
		Method:  "debug_traceTransaction",
		Params:  []interface{}{hash.String()},
	}

	body, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("marshaling params: %v", err)
	}

	req, err := http.NewRequest("POST", nodeURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request status: %d", resp.StatusCode)
	}
	f, _ := os.Create(fmt.Sprintf("/tmp/tx-%s", hash))
	io.Copy(f, resp.Body)
	f.Close()
	return nil
}
