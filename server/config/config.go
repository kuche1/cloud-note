package config

import "time"

const CertValidity = time.Hour * 24 * 365

// Lowering this will make previously created notes with
// names longer than this inaccessible
const NoteNameMaxLength uint64 = 512

// A note's content is read in one go form the filesystem (as to make it
// impossible for malicious clients to keep the filesystem locked
// almost indefinetely by lowering their download speed to the bare
// minimum), and then the read content is sent.
// Rising this to too high of a value may dramatically
// increase RAM usage.
const NoteContentsMaxLength uint64 = 1024 * 1024 * 2
