package strategy

type HandKind struct {
	ja string
	id int
}

var (
	HAND_KIND_ROCK     = HandKind{ja: "グー", id: 0}
	HAND_KIND_SCISSORS = HandKind{ja: "チョキ", id: 1}
	HAND_KIND_PAPER    = HandKind{ja: "パー", id: 2}
)

func (h *HandKind) GetJa() string {
	return h.ja
}

func (h *HandKind) GetId() int {
	return h.id
}

func GetHandKindById(id int) HandKind {
	if id == HAND_KIND_ROCK.id {
		return HAND_KIND_ROCK
	} else if id == HAND_KIND_SCISSORS.id {
		return HAND_KIND_SCISSORS
	} else {
		return HAND_KIND_PAPER
	}
}

type Hand struct {
	kind HandKind
}

func NewHand(kind HandKind) *Hand {
	return &Hand{
		kind,
	}
}

func (h *Hand) IsStrongerThan(opponent *Hand) bool {
	return h.fight(opponent) == 1
}

func (h *Hand) IsWeakerThan(opponent *Hand) bool {
	return h.fight(opponent) == -1
}

func (h *Hand) fight(opponent *Hand) int {
	if h.kind.id == opponent.kind.id {
		return 0
	} else if (h.kind.id+1)%3 == opponent.kind.id {
		return 1
	} else {
		return -1
	}
}
