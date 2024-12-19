from enum import Enum

class Turn(Enum):
    Left = 1
    Right = 2
    Reverse = 3


class Direction(Enum):
    North = 0
    East = 1
    South = 2
    West = 3

    def __lt__(self, other):
        if self.__class__ is other.__class__:
            return self.value < other.value
        return NotImplemented

    def turn(self, turn: Turn | None):
        if turn is None:
            return self

        if turn == Turn.Left:
            next_dir = self.value - 1
        elif turn == Turn.Right:
            next_dir = self.value + 1
        else:
            next_dir = self.value + 2

        if next_dir < 0:
            next_dir = next_dir + 4
        if next_dir > 3:
            next_dir = next_dir % 4

        return Direction(next_dir)


class MazeSpace(Enum):
    Empty = 0
    Wall = 1

class MazeGrid:
    def __init__(self, filename):
        self._grid: list[list[MazeSpace]] = []
        self._location: tuple[int, int] = (0, 0)
        self._facing = Direction.East
        self._goal = None
        self._walls = set()
        self._spaces = set()
        with open(filename) as f:
            for line_idx, line in enumerate(f):
                row: list[MazeSpace] = []
                for char_idx, char in enumerate(line.strip()):
                    if char == '#':
                        row.append(MazeSpace.Wall)
                        self._walls.add((line_idx, char_idx))
                    else:
                        if char == 'S':
                            self._location = line_idx, char_idx
                        elif char == 'E':
                            self._goal = line_idx, char_idx
                        row.append(MazeSpace.Empty)
                        self._spaces.add((line_idx, char_idx))
                self._grid.append(row)
            pass
        self._width = len(self._grid[0])
        self._height = len(self._grid)

    @property
    def width(self):
        return self._width

    @property
    def height(self):
        return self._height

    @property
    def walls(self):
        return self._walls

    @property
    def location(self):
        return self._location

    @property
    def grid(self):
        return self._grid

    @property
    def goal(self):
        return self._goal

    @property
    def facing(self):
        return self._facing

    @property
    def spaces(self):
        return self._spaces

    def next_location(self, direction: Direction, location: None | tuple[int, int]):
        if location is None:
            location = self._location
        if direction == Direction.North:
            return location[0] - 1, location[1]
        elif direction == Direction.East:
            return location[0], location[1] + 1
        elif direction == Direction.South:
            return location[0] + 1, location[1]
        else:
            return location[0], location[1] - 1

    def next_space(self, direction: Direction, location: None | tuple[int, int]):
        if location is None:
            location = self._location
        location = self.next_location(direction, location)
        return self._grid[location[0]][location[1]]
