package utils

type GameOverData interface {
    GetSanity() int
    GetHope() int
    GetMoney() int
    GetResumeCount() int
    GetRejections() int
    GetGhostedCount() int
}