package lib

const QuicProto = "cloud-note"

// just pick a random number and stick with it
const BytelikeACK uint8 = 202

// these 2 need to be different values (obviously)
// (but they don't have to be different from `BytelikeACK`)
const BytelikeOk uint8 = 203
const BytelikeNotOk uint8 = 204
