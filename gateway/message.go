package gateway

type MessageFunc func(*Message) *Message

type Message struct {
	Content  string
	Template string
	Data     map[string]interface{}
}

func NewMessage(content string) *Message {
	return &Message{
		Content: content,
	}
}

func (m *Message) GetContent() string {
	return m.Content
}

func (m *Message) SetContent(content string) *Message {
	m.Content = content

	return m
}

func (m *Message) SetTemplate(template string) *Message {
	m.Template = template

	return m
}

func (m *Message) SetData(data map[string]interface{}) *Message {
	m.Data = data

	return m
}

func (m *Message) GetTemplate() string {
	return m.Template
}

func (m *Message) GetData() map[string]interface{} {
	return m.Data
}
