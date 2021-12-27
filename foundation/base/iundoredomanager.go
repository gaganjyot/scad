package base

type IUndoRedoManager interface {
	CanUndo()
	CanRedo()
	Undo()
	Redo()
	ClearStack()
}
