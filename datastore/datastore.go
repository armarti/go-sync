package datastore

// Datastore abstracts over the underlying datastore.
type Datastore interface {
	// Insert a new sync entity.
	InsertSyncEntity(entity *SyncEntity) error
	// Insert a series of sync entities in a write transaction.
	InsertSyncEntitiesWithServerTags(entities []*SyncEntity) error
	// Update an existing sync entity.
	UpdateSyncEntity(entity *SyncEntity) (conflict bool, delete bool, err error)
	// Get updates for a specific type which are modified after the time of
	// client token for a given client. Besides the array of sync entities, a
	// boolean value indicating whether there are more updates to query in the
	// next batch is returned.
	GetUpdatesForType(dataType int, clientToken int64, fetchFolders bool, clientID string, maxSize int64) (bool, []SyncEntity, error)
	// Check if a server-defined unique tag is in the datastore.
	HasServerDefinedUniqueTag(clientID string, tag string) (bool, error)
	// Get the count of sync items for a client.
	GetClientItemCount(clientID string) (int, error)
	// Update the count of sync items for a client.
	UpdateClientItemCount(clientID string, count int) error
}
