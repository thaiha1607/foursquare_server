// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thaiha1607/foursquare_server/ent/conversation"
	"github.com/thaiha1607/foursquare_server/ent/message"
	"github.com/thaiha1607/foursquare_server/ent/user"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ConversationID holds the value of the "conversation_id" field.
	ConversationID uuid.UUID `json:"conversation_id,omitempty"`
	// SenderID holds the value of the "sender_id" field.
	SenderID uuid.UUID `json:"sender_id,omitempty"`
	// Type holds the value of the "type" field.
	Type message.Type `json:"type,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// IsRead holds the value of the "is_read" field.
	IsRead bool `json:"is_read,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges        MessageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// Conversation holds the value of the conversation edge.
	Conversation *Conversation `json:"conversation,omitempty"`
	// Sender holds the value of the sender edge.
	Sender *User `json:"sender,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ConversationOrErr returns the Conversation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) ConversationOrErr() (*Conversation, error) {
	if e.Conversation != nil {
		return e.Conversation, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: conversation.Label}
	}
	return nil, &NotLoadedError{edge: "conversation"}
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) SenderOrErr() (*User, error) {
	if e.Sender != nil {
		return e.Sender, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldIsRead:
			values[i] = new(sql.NullBool)
		case message.FieldType, message.FieldContent:
			values[i] = new(sql.NullString)
		case message.FieldCreatedAt, message.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case message.FieldID, message.FieldConversationID, message.FieldSenderID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case message.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case message.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case message.FieldConversationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field conversation_id", values[i])
			} else if value != nil {
				m.ConversationID = *value
			}
		case message.FieldSenderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field sender_id", values[i])
			} else if value != nil {
				m.SenderID = *value
			}
		case message.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = message.Type(value.String)
			}
		case message.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				m.Content = value.String
			}
		case message.FieldIsRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_read", values[i])
			} else if value.Valid {
				m.IsRead = value.Bool
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Message.
// This includes values selected through modifiers, order, etc.
func (m *Message) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryConversation queries the "conversation" edge of the Message entity.
func (m *Message) QueryConversation() *ConversationQuery {
	return NewMessageClient(m.config).QueryConversation(m)
}

// QuerySender queries the "sender" edge of the Message entity.
func (m *Message) QuerySender() *UserQuery {
	return NewMessageClient(m.config).QuerySender(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("conversation_id=")
	builder.WriteString(fmt.Sprintf("%v", m.ConversationID))
	builder.WriteString(", ")
	builder.WriteString("sender_id=")
	builder.WriteString(fmt.Sprintf("%v", m.SenderID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", m.Type))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(m.Content)
	builder.WriteString(", ")
	builder.WriteString("is_read=")
	builder.WriteString(fmt.Sprintf("%v", m.IsRead))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message
