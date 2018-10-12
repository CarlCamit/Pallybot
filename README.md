# Basic Architecture

===

- Connect to a service
- Respond to events
- Respond to messages

### Connect to a service

- Configurations necessary to connect
- Obtain credentials securely
- Disconnect from service cleanly

### Respond to events

- Event comes through
- Identify type of event
- Perform appropriate action

### Respond to messages

- Messages comes through
- Read message
- Perform appropriate action

# Basic Pipeline

===

- multiple events
- multiple handlers per service
- multiple goroutines per event handler

### Events

- Identify type of event
- Select appropriate broker
- Pass data to handler via channel

### Handler

- Read data from channel
- Identify type of event
- Perform appropriate action

### Goroutine

- Channels?
