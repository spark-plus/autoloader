package connections

import (
	"encoding/json"
	"time"
)

type ConnectionData struct {
	Description string    `xorm:"description"`
	Host        string    `xorm:"host"`
	Port        int       `xorm:"port"`
	Schema      string    `xorm:"schema"`
	Login       string    `xorm:"login"`
	Password    string    `xorm:"password"`
	Extra       string    `xorm:"extra"`
	CreatedAt   time.Time `xorm:"created_at"`
	UpdatedAt   time.Time `xorm:"updated_at"`
}

type Connection struct {
	ID             int64          `xorm:"'id' pk autoincr"`
	ConnectionID   string         `xorm:"connection_id"`
	ConnectionType ConnectionType `xorm:"connection_type"`
	ConnectionRef  ConnectionRef  `xorm:"json 'connection_type'"`
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
