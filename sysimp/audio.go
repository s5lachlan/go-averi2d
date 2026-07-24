package sysimp

import rl "github.com/gen2brain/raylib-go/raylib"

var AttachAudioMixedProcessor = rl.AttachAudioMixedProcessor
var CloseAudioDevice = rl.CloseAudioDevice
var DetachAudioMixedProcessor = rl.DetachAudioMixedProcessor
var InitAudioDevice = rl.InitAudioDevice
var IsAudioDeviceReady = rl.IsAudioDeviceReady
var IsAudioStreamPlaying = rl.IsAudioStreamPlaying
var IsAudioStreamProcessed = rl.IsAudioStreamProcessed
var IsAudioStreamValid = rl.IsAudioStreamValid
var PauseAudioStream = rl.PauseAudioStream
var PlayAudioStream = rl.PlayAudioStream
var ResumeAudioStream = rl.ResumeAudioStream
var SetAudioStreamBufferSizeDefault = rl.SetAudioStreamBufferSizeDefault
var SetAudioStreamCallback = rl.SetAudioStreamCallback
var SetAudioStreamPan = rl.SetAudioStreamPan
var SetAudioStreamPitch = rl.SetAudioStreamPitch
var SetAudioStreamVolume = rl.SetAudioStreamVolume
var StopAudioStream = rl.StopAudioStream
var UnloadAudioStream = rl.UnloadAudioStream
var LoadAudioStream = rl.LoadAudioStream

var GetMusicTimeLength = rl.GetMusicTimeLength
var GetMusicTimePlayed = rl.GetMusicTimePlayed
var IsMusicStreamPlaying = rl.IsMusicStreamPlaying
var IsMusicValid = rl.IsMusicValid
var PauseMusicStream = rl.PauseMusicStream
var PlayMusicStream = rl.PlayMusicStream
var ResumeMusicStream = rl.ResumeMusicStream
var SeekMusicStream = rl.SeekMusicStream
var SetMusicPan = rl.SetMusicPan
var SetMusicPitch = rl.SetMusicPitch
var SetMusicVolume = rl.SetMusicVolume
var StopMusicStream = rl.StopMusicStream
var UnloadMusicStream = rl.UnloadMusicStream
var UpdateMusicStream = rl.UpdateMusicStream
var LoadMusicStream = rl.LoadMusicStream
var LoadMusicStreamFromMemory = rl.LoadMusicStreamFromMemory

type Sound = rl.Sound

var IsSoundPlaying = rl.IsSoundPlaying
var IsSoundValid = rl.IsSoundValid
var PauseSound = rl.PauseSound
var PlaySound = rl.PlaySound
var ResumeSound = rl.ResumeSound
var SetSoundPan = rl.SetSoundPan
var SetSoundPitch = rl.SetSoundPitch
var SetSoundVolume = rl.SetSoundVolume
var StopSound = rl.StopSound
var UnloadSound = rl.UnloadSound
var UnloadSoundAlias = rl.UnloadSoundAlias
var UpdateSound = rl.UpdateSound
var LoadSound = rl.LoadSound
var LoadSoundAlias = rl.LoadSoundAlias
var LoadSoundFromWave = rl.LoadSoundFromWave
var ExportWave = rl.ExportWave
var IsWaveValid = rl.IsWaveValid
var LoadWaveSamples = rl.LoadWaveSamples
var UnloadWave = rl.UnloadWave
var UnloadWaveSamples = rl.UnloadWaveSamples
var WaveCrop = rl.WaveCrop
var WaveFormat = rl.WaveFormat
var LoadWave = rl.LoadWave
var LoadWaveFromMemory = rl.LoadWaveFromMemory
var NewWave = rl.NewWave
var WaveCopy = rl.WaveCopy
