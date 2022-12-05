package main

import (
	"testing"
)

func TestWalk(t *testing.T) {
	type args struct {
		length    int
		start     *position
		direction position
	}

	tests := []struct {
		name    string
		args    args
		want    *position
		wantErr bool
		wantPos string
	}{
		{
			name: "walk to right from start pos",
			args: args{
				length:    3,
				start:     &position{X: 1, Y: 1},
				direction: right,
			},
			want:    &position{X: 2, Y: 1},
			wantErr: false,
			wantPos: "*",
		},
		{
			name: "walk to left from start pos",
			args: args{
				length:    3,
				start:     &position{X: 1, Y: 1},
				direction: left,
			},
			want:    &position{X: 0, Y: 1},
			wantErr: false,
			wantPos: "*",
		},
		{
			name: "walk to up from start pos",
			args: args{
				length:    3,
				start:     &position{X: 1, Y: 1},
				direction: up,
			},
			want:    &position{X: 1, Y: 0},
			wantErr: false,
			wantPos: "*",
		},
		{
			name: "walk to down from start pos",
			args: args{
				length:    3,
				start:     &position{X: 1, Y: 1},
				direction: down,
			},
			want:    &position{X: 1, Y: 2},
			wantErr: false,
			wantPos: "*",
		},
		{
			name: "walk to outside from edge pos",
			args: args{
				length:    3,
				start:     &position{X: 0, Y: 0},
				direction: up,
			},
			want:    &position{X: 0, Y: 0},
			wantErr: true,
			wantPos: "S",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := newWindow(tt.args.length, *tt.args.start)
			w.direction = tt.args.direction
			err := w.walk()
			if (err != nil) != tt.wantErr {
				t.Errorf("w.walk error = %v, want: %v", err, tt.wantErr)
				return
			}
			if *w.now != *tt.want {
				t.Errorf("want: %v, but got: %v", tt.want, w.now)
				return
			}
			if w.board[w.now.Y][w.now.X] != tt.wantPos {
				t.Errorf("want: *, but got: %v", w.board[w.now.Y][w.now.X])
				return
			}
		})
	}
}

func TestRight(t *testing.T) {
	type args struct {
		direction position
	}

	tests := []struct {
		name string
		args args
		want position
	}{
		{
			name: "turn right from upward direction",
			args: args{
				direction: up,
			},
			want: right,
		},
		{
			name: "turn right from rightward direction",
			args: args{
				direction: right,
			},
			want: down,
		},
		{
			name: "turn right from downward direction",
			args: args{
				direction: down,
			},
			want: left,
		},
		{
			name: "turn right from leftward direction",
			args: args{
				direction: left,
			},
			want: up,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := newWindow(3, position{1, 1})
			w.direction = tt.args.direction
			w.turnRight()
			if w.direction != tt.want {
				t.Errorf("want: %v, but got: %v", tt.want, w.direction)
				return
			}
		})
	}
}

func TestLeft(t *testing.T) {
	type args struct {
		direction position
	}

	tests := []struct {
		name string
		args args
		want position
	}{
		{
			name: "turn left from upward direction",
			args: args{
				direction: up,
			},
			want: left,
		},
		{
			name: "turn left from rightward direction",
			args: args{
				direction: left,
			},
			want: down,
		},
		{
			name: "turn left from downward direction",
			args: args{
				direction: down,
			},
			want: right,
		},
		{
			name: "turn left from leftward direction",
			args: args{
				direction: right,
			},
			want: up,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := newWindow(3, position{1, 1})
			w.direction = tt.args.direction
			w.turnLeft()
			if w.direction != tt.want {
				t.Errorf("want: %v, but got: %v", tt.want, w.direction)
				return
			}
		})
	}

}
