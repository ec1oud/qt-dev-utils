Kitchen sink for Qt-related utilities.

**qtbug-summary**: query tool for <https://bugreports.qt.io>

Uses <https://pkg.go.dev/github.com/andygrunwald/go-jira>

Bug numbers can be given on the command line or piped in:

```
$ qtbug-summary QTBUG-54321 QTBUG-12345
QTBUG-54321 Reported P4: Is it really impossible to pass a Q_GADGET by pointer ...
QTBUG-12345 Closed P4: QVariant conversion to QChar fails
```
Developed because I sometimes keep bug fixes in git branches, and then want to
recheck later whether fixes in old branches are "done" yet:

```
$ git branch | qtbug-summary
...
  QTBUG-35051 Open P3: Focus scope does not handle disabling active item
  QTBUG-35688 Open P2: QQuickTextEdit: Make it possible to change the text document
  QTBUG-40856-QTBUG-120346-etc
	QTBUG-40856 Open P2: MouseArea containsMouse flag is not reset ...
	QTBUG-120346 Closed P2:  hovered property of HoverHandler stays ...
  QTBUG-74496-revert1 Closed P2: Performance issue: rejected drag re-triggers ...
  QTBUG-75223 Closed P2: TapHandlers in flickable block flicking on touch screen
  QTBUG-77629-fix Closed P2: QML Flickable stealing touch events from PinchArea
```
I also keep test cases and notes about bugs and can check status this way:

```
$ cd ~/dev/bugs
$ ls | qtbug-summary
...
QTBUG-100068 Closed P3: Colors of some SVG images are wrong
QTBUG-100104 Open P2: Using TapHandler in a QML Dialog incorrectly propagates ...
QTBUG-100110 Closed P2: switch handle goes out of bounds when size is increased ...
...
```

