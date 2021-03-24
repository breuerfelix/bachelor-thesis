func (c *ConnectPacket) Unpack(b io.Reader) {
	c.ProtocolName, _ = decodeString(b)
	c.ProtocolVersion, _ = decodeByte(b)
	options, _ := decodeByte(b)
	c.ReservedBit = 1 & options
	c.CleanSession = 1 & (options >> 1) > 0
	c.WillFlag = 1 & (options >> 2) > 0
	c.WillQos = 3 & (options >> 3)
	c.WillRetain = 1 & (options >> 5) > 0
	c.PasswordFlag = 1 & (options >> 6) > 0
	c.UsernameFlag = 1 & (options >> 7) > 0
	c.Keepalive, _ = decodeUint16(b)
	c.ClientIdentifier, _ = decodeString(b)
	if c.WillFlag {
		c.WillTopic, _ = decodeString(b)
		c.WillMessage, _ = decodeBytes(b)
	}
	if c.UsernameFlag {
		c.Username, _ = decodeString(b)
	}
	if c.PasswordFlag {
		c.Password, _ = decodeBytes(b)
	}
}
