package main

type entry struct {
	code   uint64
	length int
}

var encodeTable = [...]entry{
	0x00:  {0b111111111111111111111000, 24},
	0x01:  {0b111111111111111111111001, 24},
	0x02:  {0b111111111111111111111010, 24},
	0x03:  {0b111111111111111111111011, 24},
	0x04:  {0b111111111111111111000000, 24},
	0x05:  {0b111111111111111111000001, 24},
	0x06:  {0b111111111111111111000010, 24},
	0x07:  {0b111111111111111111000011, 24},
	0x08:  {0b111111111111111111000100, 24},
	0x09:  {0b111111111111111111000101, 24},
	0x0a:  {0b111111111111111111000110, 24},
	0x0b:  {0b111111111111111111000111, 24},
	0x0c:  {0b111111111111111111001000, 24},
	0x0d:  {0b111111111111111111001001, 24},
	0x0e:  {0b111111111111111111001010, 24},
	0x0f:  {0b111111111111111111001011, 24},
	0x10:  {0b111111111111111111001100, 24},
	0x11:  {0b111111111111111111001101, 24},
	0x12:  {0b111111111111111111001110, 24},
	0x13:  {0b111111111111111111001111, 24},
	0x14:  {0b111111111111111111010000, 24},
	0x15:  {0b111111111111111111010001, 24},
	0x16:  {0b111111111111111111010010, 24},
	0x17:  {0b111111111111111111010011, 24},
	0x18:  {0b111111111111111111010100, 24},
	0x19:  {0b111111111111111111010101, 24},
	0x1a:  {0b111111111111111111010110, 24},
	0x1b:  {0b111111111111111111010111, 24},
	0x1c:  {0b111111111111111111011000, 24},
	0x1d:  {0b111111111111111111011001, 24},
	0x1e:  {0b111111111111111111011010, 24},
	0x1f:  {0b111111111111111111011011, 24},
	0x20:  {0b111111111111111111011100, 24},
	0x21:  {0b11111111111110, 14},
	0x22:  {0b111111111111111111011101, 24},
	0x23:  {0b01001011, 8},
	0x24:  {0b111111111111111111011110, 24},
	0x25:  {0b111110, 6},
	0x26:  {0b100001100010, 12},
	0x27:  {0b01001010010100, 14},
	0x28:  {0b1111111111111110, 16},
	0x29:  {0b111111111111110, 15},
	0x2a:  {0b111111111111111111011111, 24},
	0x2b:  {0b0100101001011, 13},
	0x2c:  {0b111111111111111110000000, 24},
	0x2d:  {0b11110, 5},
	0x2e:  {0b10001, 5},
	0x2f:  {0b1101, 4},
	0x30:  {0b101001, 6},
	0x31:  {0b01000, 5},
	0x32:  {0b01111, 5},
	0x33:  {0b011010, 6},
	0x34:  {0b011011, 6},
	0x35:  {0b1000010, 7},
	0x36:  {0b1111110, 7},
	0x37:  {0b0100110, 7},
	0x38:  {0b101100, 6},
	0x39:  {0b1000000, 7},
	0x3a:  {0b1111111101, 10},
	0x3b:  {0b111111111111111110000001, 24},
	0x3c:  {0b111111111111111110000010, 24},
	0x3d:  {0b1001111000, 10},
	0x3e:  {0b111111111111111110000011, 24},
	0x3f:  {0b0100111110, 10},
	0x40:  {0b01001010010101, 14},
	0x41:  {0b01001110, 8},
	0x42:  {0b100111101, 9},
	0x43:  {0b1001111110, 10},
	0x44:  {0b1111111110, 10},
	0x45:  {0b0100100, 7},
	0x46:  {0b1111111100, 10},
	0x47:  {0b10011110010, 11},
	0x48:  {0b01001111010, 11},
	0x49:  {0b10011110011, 11},
	0x4a:  {0b100001100011, 12},
	0x4b:  {0b100001100100, 12},
	0x4c:  {0b0100111111, 10},
	0x4d:  {0b0100101010, 10},
	0x4e:  {0b01001010011, 11},
	0x4f:  {0b10000110000, 11},
	0x50:  {0b0100101000, 10},
	0x51:  {0b100001100101, 12},
	0x52:  {0b0100111100, 10},
	0x53:  {0b1001111111, 10},
	0x54:  {0b0100101011, 10},
	0x55:  {0b10000110011, 11},
	0x56:  {0b010011110110, 12},
	0x57:  {0b11111111110, 11},
	0x58:  {0b010011110111, 12},
	0x59:  {0b111111111110, 12},
	0x5a:  {0b1111111111110, 13},
	0x5b:  {0b111111111111111110000100, 24},
	0x5c:  {0b111111111111111110000101, 24},
	0x5d:  {0b111111111111111110000110, 24},
	0x5e:  {0b111111111111111110000111, 24},
	0x5f:  {0b10011100, 8},
	0x60:  {0b111111111111111110001000, 24},
	0x61:  {0b0000, 4},
	0x62:  {0b101101, 6},
	0x63:  {0b11100, 5},
	0x64:  {0b101000, 6},
	0x65:  {0b0011, 4},
	0x66:  {0b1000001, 7},
	0x67:  {0b10010, 5},
	0x68:  {0b01100, 5},
	0x69:  {0b10111, 5},
	0x6a:  {0b10011101, 8},
	0x6b:  {0b0111011, 7},
	0x6c:  {0b01011, 5},
	0x6d:  {0b00100, 5},
	0x6e:  {0b01010, 5},
	0x6f:  {0b1100, 4},
	0x70:  {0b100110, 6},
	0x71:  {0b11111110, 8},
	0x72:  {0b11101, 5},
	0x73:  {0b10101, 5},
	0x74:  {0b0001, 4},
	0x75:  {0b00101, 5},
	0x76:  {0b10000111, 8},
	0x77:  {0b011100, 6},
	0x78:  {0b100001101, 9},
	0x79:  {0b0111010, 7},
	0x7a:  {0b100111110, 9},
	0x7b:  {0b111111111111111110001001, 24},
	0x7c:  {0b111111111111111110001010, 24},
	0x7d:  {0b111111111111111110001011, 24},
	0x7e:  {0b010010100100, 12},
	0x7f:  {0b111111111111111110001100, 24},
	0x80:  {0b111111111111111110001101, 24},
	0x81:  {0b111111111111111110001110, 24},
	0x82:  {0b111111111111111110001111, 24},
	0x83:  {0b111111111111111110010000, 24},
	0x84:  {0b111111111111111110010001, 24},
	0x85:  {0b111111111111111110010010, 24},
	0x86:  {0b111111111111111110010011, 24},
	0x87:  {0b111111111111111110010100, 24},
	0x88:  {0b111111111111111110010101, 24},
	0x89:  {0b111111111111111110010110, 24},
	0x8a:  {0b111111111111111110010111, 24},
	0x8b:  {0b111111111111111110011000, 24},
	0x8c:  {0b111111111111111110011001, 24},
	0x8d:  {0b111111111111111110011010, 24},
	0x8e:  {0b111111111111111110011011, 24},
	0x8f:  {0b111111111111111110011100, 24},
	0x90:  {0b111111111111111110011101, 24},
	0x91:  {0b111111111111111110011110, 24},
	0x92:  {0b111111111111111110011111, 24},
	0x93:  {0b111111111111111110100000, 24},
	0x94:  {0b111111111111111110100001, 24},
	0x95:  {0b111111111111111110100010, 24},
	0x96:  {0b111111111111111110100011, 24},
	0x97:  {0b111111111111111110100100, 24},
	0x98:  {0b111111111111111110100101, 24},
	0x99:  {0b111111111111111110100110, 24},
	0x9a:  {0b111111111111111110100111, 24},
	0x9b:  {0b111111111111111110101000, 24},
	0x9c:  {0b111111111111111110101001, 24},
	0x9d:  {0b111111111111111110101010, 24},
	0x9e:  {0b111111111111111110101011, 24},
	0x9f:  {0b111111111111111110101100, 24},
	0xa0:  {0b111111111111111110101101, 24},
	0xa1:  {0b111111111111111110101110, 24},
	0xa2:  {0b111111111111111110101111, 24},
	0xa3:  {0b111111111111111110110000, 24},
	0xa4:  {0b111111111111111110110001, 24},
	0xa5:  {0b111111111111111110110010, 24},
	0xa6:  {0b111111111111111110110011, 24},
	0xa7:  {0b111111111111111110110100, 24},
	0xa8:  {0b111111111111111110110101, 24},
	0xa9:  {0b111111111111111110110110, 24},
	0xaa:  {0b111111111111111110110111, 24},
	0xab:  {0b111111111111111110111000, 24},
	0xac:  {0b111111111111111110111001, 24},
	0xad:  {0b111111111111111110111010, 24},
	0xae:  {0b111111111111111110111011, 24},
	0xaf:  {0b111111111111111110111100, 24},
	0xb0:  {0b111111111111111110111101, 24},
	0xb1:  {0b111111111111111110111110, 24},
	0xb2:  {0b111111111111111110111111, 24},
	0xb3:  {0b11111111111111110000000, 23},
	0xb4:  {0b11111111111111110000001, 23},
	0xb5:  {0b11111111111111110000010, 23},
	0xb6:  {0b11111111111111110000011, 23},
	0xb7:  {0b11111111111111110000100, 23},
	0xb8:  {0b11111111111111110000101, 23},
	0xb9:  {0b11111111111111110000110, 23},
	0xba:  {0b11111111111111110000111, 23},
	0xbb:  {0b11111111111111110001000, 23},
	0xbc:  {0b11111111111111110001001, 23},
	0xbd:  {0b11111111111111110001010, 23},
	0xbe:  {0b11111111111111110001011, 23},
	0xbf:  {0b11111111111111110001100, 23},
	0xc0:  {0b11111111111111110001101, 23},
	0xc1:  {0b11111111111111110001110, 23},
	0xc2:  {0b11111111111111110001111, 23},
	0xc3:  {0b11111111111111110010000, 23},
	0xc4:  {0b11111111111111110010001, 23},
	0xc5:  {0b11111111111111110010010, 23},
	0xc6:  {0b11111111111111110010011, 23},
	0xc7:  {0b11111111111111110010100, 23},
	0xc8:  {0b11111111111111110010101, 23},
	0xc9:  {0b11111111111111110010110, 23},
	0xca:  {0b11111111111111110010111, 23},
	0xcb:  {0b11111111111111110011000, 23},
	0xcc:  {0b11111111111111110011001, 23},
	0xcd:  {0b11111111111111110011010, 23},
	0xce:  {0b11111111111111110011011, 23},
	0xcf:  {0b11111111111111110011100, 23},
	0xd0:  {0b11111111111111110011101, 23},
	0xd1:  {0b11111111111111110011110, 23},
	0xd2:  {0b11111111111111110011111, 23},
	0xd3:  {0b11111111111111110100000, 23},
	0xd4:  {0b11111111111111110100001, 23},
	0xd5:  {0b11111111111111110100010, 23},
	0xd6:  {0b11111111111111110100011, 23},
	0xd7:  {0b11111111111111110100100, 23},
	0xd8:  {0b11111111111111110100101, 23},
	0xd9:  {0b11111111111111110100110, 23},
	0xda:  {0b11111111111111110100111, 23},
	0xdb:  {0b11111111111111110101000, 23},
	0xdc:  {0b11111111111111110101001, 23},
	0xdd:  {0b11111111111111110101010, 23},
	0xde:  {0b11111111111111110101011, 23},
	0xdf:  {0b11111111111111110101100, 23},
	0xe0:  {0b11111111111111110101101, 23},
	0xe1:  {0b11111111111111110101110, 23},
	0xe2:  {0b11111111111111110101111, 23},
	0xe3:  {0b11111111111111110110000, 23},
	0xe4:  {0b11111111111111110110001, 23},
	0xe5:  {0b11111111111111110110010, 23},
	0xe6:  {0b11111111111111110110011, 23},
	0xe7:  {0b11111111111111110110100, 23},
	0xe8:  {0b11111111111111110110101, 23},
	0xe9:  {0b11111111111111110110110, 23},
	0xea:  {0b11111111111111110110111, 23},
	0xeb:  {0b11111111111111110111000, 23},
	0xec:  {0b11111111111111110111001, 23},
	0xed:  {0b11111111111111110111010, 23},
	0xee:  {0b11111111111111110111011, 23},
	0xef:  {0b11111111111111110111100, 23},
	0xf0:  {0b11111111111111110111101, 23},
	0xf1:  {0b11111111111111110111110, 23},
	0xf2:  {0b11111111111111110111111, 23},
	0xf3:  {0b11111111111111111110000, 23},
	0xf4:  {0b11111111111111111110001, 23},
	0xf5:  {0b11111111111111111110010, 23},
	0xf6:  {0b11111111111111111110011, 23},
	0xf7:  {0b11111111111111111110100, 23},
	0xf8:  {0b11111111111111111110101, 23},
	0xf9:  {0b11111111111111111110110, 23},
	0xfa:  {0b11111111111111111110111, 23},
	0xfb:  {0b11111111111111111111000, 23},
	0xfc:  {0b11111111111111111111001, 23},
	0xfd:  {0b11111111111111111111010, 23},
	0xfe:  {0b11111111111111111111011, 23},
	0xff:  {0b11111111111111111111110, 23},
	0x100: {0b11111111111111111111111, 23},
}

type node struct {
	value    int
	children [2]*node
}

var decodeTree = &node{children: [2]*node{
	{children: [2]*node{
		{children: [2]*node{
			{children: [2]*node{
				{value: 0x61},
				{value: 0x74},
			}},
			{children: [2]*node{
				{children: [2]*node{
					{value: 0x6d},
					{value: 0x75},
				}},
				{value: 0x65},
			}},
		}},
		{children: [2]*node{
			{children: [2]*node{
				{children: [2]*node{
					{value: 0x31},
					{children: [2]*node{
						{children: [2]*node{
							{value: 0x45},
							{children: [2]*node{
								{children: [2]*node{
									{children: [2]*node{
										{value: 0x50},
										{children: [2]*node{
											{children: [2]*node{
												{value: 0x7e},
												{children: [2]*node{
													{children: [2]*node{
														{value: 0x27},
														{value: 0x40},
													}},
													{value: 0x2b},
												}},
											}},
											{value: 0x4e},
										}},
									}},
									{children: [2]*node{
										{value: 0x4d},
										{value: 0x54},
									}},
								}},
								{value: 0x23},
							}},
						}},
						{children: [2]*node{
							{value: 0x37},
							{children: [2]*node{
								{value: 0x41},
								{children: [2]*node{
									{children: [2]*node{
										{value: 0x52},
										{children: [2]*node{
											{value: 0x48},
											{children: [2]*node{
												{value: 0x56},
												{value: 0x58},
											}},
										}},
									}},
									{children: [2]*node{
										{value: 0x3f},
										{value: 0x4c},
									}},
								}},
							}},
						}},
					}},
				}},
				{children: [2]*node{
					{value: 0x6e},
					{value: 0x6c},
				}},
			}},
			{children: [2]*node{
				{children: [2]*node{
					{value: 0x68},
					{children: [2]*node{
						{value: 0x33},
						{value: 0x34},
					}},
				}},
				{children: [2]*node{
					{children: [2]*node{
						{value: 0x77},
						{children: [2]*node{
							{value: 0x79},
							{value: 0x6b},
						}},
					}},
					{value: 0x32},
				}},
			}},
		}},
	}},
	{children: [2]*node{
		{children: [2]*node{
			{children: [2]*node{
				{children: [2]*node{
					{children: [2]*node{
						{children: [2]*node{
							{value: 0x39},
							{value: 0x66},
						}},
						{children: [2]*node{
							{value: 0x35},
							{children: [2]*node{
								{children: [2]*node{
									{children: [2]*node{
										{children: [2]*node{
											{value: 0x4f},
											{children: [2]*node{
												{value: 0x26},
												{value: 0x4a},
											}},
										}},
										{children: [2]*node{
											{children: [2]*node{
												{value: 0x4b},
												{value: 0x51},
											}},
											{value: 0x55},
										}},
									}},
									{value: 0x78},
								}},
								{value: 0x76},
							}},
						}},
					}},
					{value: 0x2e},
				}},
				{children: [2]*node{
					{value: 0x67},
					{children: [2]*node{
						{value: 0x70},
						{children: [2]*node{
							{children: [2]*node{
								{value: 0x5f},
								{value: 0x6a},
							}},
							{children: [2]*node{
								{children: [2]*node{
									{children: [2]*node{
										{value: 0x3d},
										{children: [2]*node{
											{value: 0x47},
											{value: 0x49},
										}},
									}},
									{value: 0x42},
								}},
								{children: [2]*node{
									{value: 0x7a},
									{children: [2]*node{
										{value: 0x43},
										{value: 0x53},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
			{children: [2]*node{
				{children: [2]*node{
					{children: [2]*node{
						{value: 0x64},
						{value: 0x30},
					}},
					{value: 0x73},
				}},
				{children: [2]*node{
					{children: [2]*node{
						{value: 0x38},
						{value: 0x62},
					}},
					{value: 0x69},
				}},
			}},
		}},
		{children: [2]*node{
			{children: [2]*node{
				{value: 0x6f},
				{value: 0x2f},
			}},
			{children: [2]*node{
				{children: [2]*node{
					{value: 0x63},
					{value: 0x72},
				}},
				{children: [2]*node{
					{value: 0x2d},
					{children: [2]*node{
						{value: 0x25},
						{children: [2]*node{
							{value: 0x36},
							{children: [2]*node{
								{value: 0x71},
								{children: [2]*node{
									{children: [2]*node{
										{value: 0x46},
										{value: 0x3a},
									}},
									{children: [2]*node{
										{value: 0x44},
										{children: [2]*node{
											{value: 0x57},
											{children: [2]*node{
												{value: 0x59},
												{children: [2]*node{
													{value: 0x5a},
													{children: [2]*node{
														{value: 0x21},
														{children: [2]*node{
															{value: 0x29},
															{children: [2]*node{
																{value: 0x28},
																{children: [2]*node{
																	{children: [2]*node{
																		{children: [2]*node{
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xb3},
																							{value: 0xb4},
																						}},
																						{children: [2]*node{
																							{value: 0xb5},
																							{value: 0xb6},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xb7},
																							{value: 0xb8},
																						}},
																						{children: [2]*node{
																							{value: 0xb9},
																							{value: 0xba},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xbb},
																							{value: 0xbc},
																						}},
																						{children: [2]*node{
																							{value: 0xbd},
																							{value: 0xbe},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xbf},
																							{value: 0xc0},
																						}},
																						{children: [2]*node{
																							{value: 0xc1},
																							{value: 0xc2},
																						}},
																					}},
																				}},
																			}},
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xc3},
																							{value: 0xc4},
																						}},
																						{children: [2]*node{
																							{value: 0xc5},
																							{value: 0xc6},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xc7},
																							{value: 0xc8},
																						}},
																						{children: [2]*node{
																							{value: 0xc9},
																							{value: 0xca},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xcb},
																							{value: 0xcc},
																						}},
																						{children: [2]*node{
																							{value: 0xcd},
																							{value: 0xce},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xcf},
																							{value: 0xd0},
																						}},
																						{children: [2]*node{
																							{value: 0xd1},
																							{value: 0xd2},
																						}},
																					}},
																				}},
																			}},
																		}},
																		{children: [2]*node{
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xd3},
																							{value: 0xd4},
																						}},
																						{children: [2]*node{
																							{value: 0xd5},
																							{value: 0xd6},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xd7},
																							{value: 0xd8},
																						}},
																						{children: [2]*node{
																							{value: 0xd9},
																							{value: 0xda},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xdb},
																							{value: 0xdc},
																						}},
																						{children: [2]*node{
																							{value: 0xdd},
																							{value: 0xde},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xdf},
																							{value: 0xe0},
																						}},
																						{children: [2]*node{
																							{value: 0xe1},
																							{value: 0xe2},
																						}},
																					}},
																				}},
																			}},
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xe3},
																							{value: 0xe4},
																						}},
																						{children: [2]*node{
																							{value: 0xe5},
																							{value: 0xe6},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xe7},
																							{value: 0xe8},
																						}},
																						{children: [2]*node{
																							{value: 0xe9},
																							{value: 0xea},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xeb},
																							{value: 0xec},
																						}},
																						{children: [2]*node{
																							{value: 0xed},
																							{value: 0xee},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xef},
																							{value: 0xf0},
																						}},
																						{children: [2]*node{
																							{value: 0xf1},
																							{value: 0xf2},
																						}},
																					}},
																				}},
																			}},
																		}},
																	}},
																	{children: [2]*node{
																		{children: [2]*node{
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x2c},
																								{value: 0x3b},
																							}},
																							{children: [2]*node{
																								{value: 0x3c},
																								{value: 0x3e},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x5b},
																								{value: 0x5c},
																							}},
																							{children: [2]*node{
																								{value: 0x5d},
																								{value: 0x5e},
																							}},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x60},
																								{value: 0x7b},
																							}},
																							{children: [2]*node{
																								{value: 0x7c},
																								{value: 0x7d},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x7f},
																								{value: 0x80},
																							}},
																							{children: [2]*node{
																								{value: 0x81},
																								{value: 0x82},
																							}},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x83},
																								{value: 0x84},
																							}},
																							{children: [2]*node{
																								{value: 0x85},
																								{value: 0x86},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x87},
																								{value: 0x88},
																							}},
																							{children: [2]*node{
																								{value: 0x89},
																								{value: 0x8a},
																							}},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x8b},
																								{value: 0x8c},
																							}},
																							{children: [2]*node{
																								{value: 0x8d},
																								{value: 0x8e},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x8f},
																								{value: 0x90},
																							}},
																							{children: [2]*node{
																								{value: 0x91},
																								{value: 0x92},
																							}},
																						}},
																					}},
																				}},
																			}},
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x93},
																								{value: 0x94},
																							}},
																							{children: [2]*node{
																								{value: 0x95},
																								{value: 0x96},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x97},
																								{value: 0x98},
																							}},
																							{children: [2]*node{
																								{value: 0x99},
																								{value: 0x9a},
																							}},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x9b},
																								{value: 0x9c},
																							}},
																							{children: [2]*node{
																								{value: 0x9d},
																								{value: 0x9e},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x9f},
																								{value: 0xa0},
																							}},
																							{children: [2]*node{
																								{value: 0xa1},
																								{value: 0xa2},
																							}},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0xa3},
																								{value: 0xa4},
																							}},
																							{children: [2]*node{
																								{value: 0xa5},
																								{value: 0xa6},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0xa7},
																								{value: 0xa8},
																							}},
																							{children: [2]*node{
																								{value: 0xa9},
																								{value: 0xaa},
																							}},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0xab},
																								{value: 0xac},
																							}},
																							{children: [2]*node{
																								{value: 0xad},
																								{value: 0xae},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0xaf},
																								{value: 0xb0},
																							}},
																							{children: [2]*node{
																								{value: 0xb1},
																								{value: 0xb2},
																							}},
																						}},
																					}},
																				}},
																			}},
																		}},
																		{children: [2]*node{
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x04},
																								{value: 0x05},
																							}},
																							{children: [2]*node{
																								{value: 0x06},
																								{value: 0x07},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x08},
																								{value: 0x09},
																							}},
																							{children: [2]*node{
																								{value: 0x0a},
																								{value: 0x0b},
																							}},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x0c},
																								{value: 0x0d},
																							}},
																							{children: [2]*node{
																								{value: 0x0e},
																								{value: 0x0f},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x10},
																								{value: 0x11},
																							}},
																							{children: [2]*node{
																								{value: 0x12},
																								{value: 0x13},
																							}},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x14},
																								{value: 0x15},
																							}},
																							{children: [2]*node{
																								{value: 0x16},
																								{value: 0x17},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x18},
																								{value: 0x19},
																							}},
																							{children: [2]*node{
																								{value: 0x1a},
																								{value: 0x1b},
																							}},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x1c},
																								{value: 0x1d},
																							}},
																							{children: [2]*node{
																								{value: 0x1e},
																								{value: 0x1f},
																							}},
																						}},
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x20},
																								{value: 0x22},
																							}},
																							{children: [2]*node{
																								{value: 0x24},
																								{value: 0x2a},
																							}},
																						}},
																					}},
																				}},
																			}},
																			{children: [2]*node{
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xf3},
																							{value: 0xf4},
																						}},
																						{children: [2]*node{
																							{value: 0xf5},
																							{value: 0xf6},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xf7},
																							{value: 0xf8},
																						}},
																						{children: [2]*node{
																							{value: 0xf9},
																							{value: 0xfa},
																						}},
																					}},
																				}},
																				{children: [2]*node{
																					{children: [2]*node{
																						{children: [2]*node{
																							{value: 0xfb},
																							{value: 0xfc},
																						}},
																						{children: [2]*node{
																							{value: 0xfd},
																							{value: 0xfe},
																						}},
																					}},
																					{children: [2]*node{
																						{children: [2]*node{
																							{children: [2]*node{
																								{value: 0x00},
																								{value: 0x01},
																							}},
																							{children: [2]*node{
																								{value: 0x02},
																								{value: 0x03},
																							}},
																						}},
																						{children: [2]*node{
																							{value: 0xff},
																							{value: 0x100},
																						}},
																					}},
																				}},
																			}},
																		}},
																	}},
																}},
															}},
														}},
													}},
												}},
											}},
										}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}},
}}
