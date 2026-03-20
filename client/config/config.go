package config

// These values can be lowered if you suspect that the server wants
// to flood your RAM
const NoteNameMaxLength uint64 = 512 * 2
const NoteContentsMaxLength uint64 = 1024 * 1024 * 2 * 2
const NumberOfNotesMaxLength uint64 = 512
