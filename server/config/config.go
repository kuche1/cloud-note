package config

import "time"

const CertValidity = time.Hour * 24 * 365

// TODO:
// I don't feel great passing the values below everywhere
// I need to make some other functions that take them into
// account
// Same goes for the client

// IMPROVE000: Make message readable for the end user
// Lowering this will make previously created notes with
// names longer than this inaccessible
const NoteNameMaxLength uint64 = 512

// IMPROVE000: Make message readable for the end user
// A note's content is read in one go (as to make it
// impossible for malicious clients to keep the filesystem locked
// almost indefinetely by lowering their download speed to the bare
// minimum), and then the read content is sent.
// Rising this to too high of a value may dramatically
// increase RAM usage.
const NoteContentsMaxLength uint64 = 1024 * 1024 * 2
