package cmd

import (
	"testing"

	"github.com/go-zookeeper/zk"
)

type stubZKPathClient struct {
	existing map[string]bool
	created  []string
}

func (s *stubZKPathClient) Exists(path string) (bool, *zk.Stat, error) {
	return s.existing[path], nil, nil
}

func (s *stubZKPathClient) Create(path string, data []byte, flags int32, acl []zk.ACL) (string, error) {
	s.existing[path] = true
	s.created = append(s.created, path)
	return path, nil
}

func TestEnsureParentExistsCreatesIntermediatePaths(t *testing.T) {
	stub := &stubZKPathClient{
		existing: map[string]bool{
			"/config": true,
		},
	}

	ensureParentExists(stub, "/config/product/service/feature")

	expected := []string{"/config/product", "/config/product/service"}
	if len(stub.created) != len(expected) {
		t.Fatalf("expected %d created paths, got %d (%v)", len(expected), len(stub.created), stub.created)
	}

	for i, want := range expected {
		if stub.created[i] != want {
			t.Fatalf("created path %d: want %q, got %q", i, want, stub.created[i])
		}
	}
}
