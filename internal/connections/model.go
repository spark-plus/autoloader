package connections

import (
	"encoding/json"
	"time"
)

type ConnectionData struct {
	Description string    `json:"description"`
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	Schema      string    `json:"schema"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Extra       string    `json:"extra"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Connection struct {
	ID             int64          `xorm:"'id' pk autoincr"`
	ConnectionID   string         `xorm:"connection_id"`
	ConnectionType ConnectionType `xorm:"connection_type"`
	ConnectionRef  ConnectionRef  `xorm:"json 'connection_ref'"`
	Description    string         `xorm:"description"`
	CreatedAt      time.Time      `xorm:"created_at"`
	UpdatedAt      time.Time      `xorm:"updated_at"`
}

type ConnectionRef struct {
	Location     LocationType `json:"location_type"`
	LocationPath string       `json:"location_path"`
}

type JSONConnectionRef ConnectionRef

func (ja *JSONConnectionRef) FromDB(data []byte) error {
	return json.Unmarshal(data, ja)
}

func (ja *JSONConnectionRef) ToDB() ([]byte, error) {
	return json.Marshal(ja)
}
