package secret

import (
	"time"

	"github.com/grafana/grafana/pkg/registry/apis/secret/encryption"
)

// EncryptionDataKey does not have a mirrored K8s resource
type EncryptionDataKey struct {
	UID           string                `xorm:"pk 'uid'"`
	Active        bool                  `xorm:"active"`
	Namespace     string                `xorm:"namespace"`
	Label         string                `xorm:"label"`
	Scope         string                `xorm:"scope"`
	Provider      encryption.ProviderID `xorm:"provider"`
	EncryptedData []byte                `xorm:"encrypted_data"`
	Created       time.Time             `xorm:"created"`
	Updated       time.Time             `xorm:"updated"`
}
