package post

type PostStatu int

const (
	PostStatuNotExsit PostStatu = -1
	PostStatuPublic   PostStatu = 0
	PostStatuPrivate  PostStatu = 1
	PostStatuDraft    PostStatu = 2
)

func IToPostStatu(i int) PostStatu {
	return itoPostStatu(i, PostStatuPublic)
}

func itoPostStatu(i int, d PostStatu) PostStatu {
	if i == int(PostStatuDraft) {
		return PostStatuDraft
	} else if i == int(PostStatuPrivate) {
		return PostStatuPrivate
	} else if i == int(PostStatuDraft) {
		return PostStatuDraft
	}

	return d
}
