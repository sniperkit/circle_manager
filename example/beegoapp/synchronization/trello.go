package synchronization

import (
	trello "github.com/VojtechVitek/go-trello"
	"github.com/jungju/circle_manager/example/beegoapp/models"
	"github.com/jungju/circle_manager/modules"
	"github.com/sirupsen/logrus"
)

func SyncTello(user string, key string, token string) error {
	myTrello, err := trello.NewAuthClient(key, &token)
	if err != nil {
		return err
	}

	member, err := myTrello.Member(user)
	if err != nil {
		return err
	}

	trelloBoards, err := member.Boards()
	if err != nil {
		return err
	}

	cards := []trello.Card{}
	lists := []trello.List{}
	mapBoard := map[string]trello.Board{}
	for _, trelloBoard := range trelloBoards {
		mapBoard[trelloBoard.Id] = trelloBoard
		cs, err := trelloBoard.Cards()
		if err != nil {
			return err
		}
		ls, err := trelloBoard.Lists()
		if err != nil {
			return err
		}

		cards = append(cards, cs...)
		lists = append(lists, ls...)
	}

	mapList := map[string]trello.List{}
	for _, list := range lists {
		mapList[list.Id] = list
	}

	gotTodos := []models.Todo{}
	if err := modules.GetItems(&gotTodos, nil); err != nil {
		return err
	}

	mapGotTodos := map[string]models.Todo{}
	mapExistsTodos := map[string]bool{}
	for _, gotTodo := range gotTodos {
		mapGotTodos[gotTodo.CardID] = gotTodo
		mapExistsTodos[gotTodo.CardID] = false
	}

	for _, card := range cards {
		listName := "unknown"
		if _, ok := mapList[card.IdList]; ok {
			listName = mapList[card.IdList].Name
		}
		boardID := ""
		boardName := ""
		if board, ok := mapBoard[card.IdBoard]; ok {
			boardID = board.Id
			boardName = board.Name
		}

		if _, ok := mapExistsTodos[card.Id]; ok {
			mapExistsTodos[card.Id] = true
		}

		updateItem := &models.Todo{
			Name:      card.Name,
			ListID:    card.IdList,
			ListName:  listName,
			CardID:    card.Id,
			BoardID:   boardID,
			BoardName: boardName,
		}

		if getItem, ok := mapGotTodos[card.Id]; !ok {
			if err := modules.CreateItem(updateItem); err != nil {
				return err
			}
		} else {
			updateItem.ID = getItem.ID
			updateItem.CreatedAt = getItem.CreatedAt

			if !IsEqual(updateItem, getItem) {
				if err := modules.SaveItem(updateItem); err != nil {
					return err
				}
			}
		}
	}

	for uuidKey, exists := range mapExistsTodos {
		if !exists {
			if err := modules.DeleteItemByColName("todos", "card_id", uuidKey); err != nil {
				logrus.WithError(err).Error()
			}
		}
	}

	return nil
}
