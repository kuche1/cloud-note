package config

import "math"

// These values can be lowered if you suspect that the server wants
// to flood your RAM
// TODO: Add some better defaults
const NoteNameMaxLength uint64 = math.MaxUint64
const NoteContentsMaxLength uint64 = math.MaxUint64
const NumberOfNotesMaxLength uint64 = math.MaxUint64
