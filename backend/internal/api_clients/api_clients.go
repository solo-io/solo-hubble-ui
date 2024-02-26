package api_clients

import (
	"context"

	"github.com/cilium/hubble-ui/backend/internal/ns_watcher"
	"github.com/cilium/hubble-ui/backend/internal/relay_client"
	"github.com/cilium/hubble-ui/backend/soloio/storage/remote"
)

type APIClientsInterface interface {
	RelayClient() relay_client.RelayClientInterface
	NSWatcher(context.Context, ns_watcher.NSWatcherOptions) (ns_watcher.NSWatcherInterface, error)
	SnapshotReader() remote.Reader
}
