#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define BOARD_SIZE 8

#define QUEEN_WEIGHT 9
#define ROOK_WEIGHT 5
#define BISHOP_WEIGHT 3
#define KNIGHT_WEIGHT 3
#define PAWN_WEIGHT 1
#define KING_WEIGHT 0

int main (int argc, char * argv[]) {
    int score_white = 0, score_black = 0;
    char current;
    for (int i = 0; i < BOARD_SIZE; i++) {
        for (int j = 0; j < BOARD_SIZE + 1; j++) { // '\n'
            scanf ("%c", &current);
            if ('Q' == current) {
                score_white += QUEEN_WEIGHT;
            } else if ('R' == current) {
                score_white += ROOK_WEIGHT;
            } else if ('B' == current) {
                score_white += BISHOP_WEIGHT;
            } else if ('N' == current) {
                score_white += KNIGHT_WEIGHT;
            } else if ('P' == current) {
                score_white += PAWN_WEIGHT;
            } else if ('K' == current) {
                score_white += KING_WEIGHT;
            } else if ('q' == current) {
                score_black += QUEEN_WEIGHT;
            } else if ('r' == current) {
                score_black += ROOK_WEIGHT;
            } else if ('b' == current) {
                score_black += BISHOP_WEIGHT;
            } else if ('n' == current) {
                score_black += KNIGHT_WEIGHT;
            } else if ('p' == current) {
                score_black += PAWN_WEIGHT;
            } else if ('k' == current) {
                score_black += KING_WEIGHT;
            }
        }
    }

    if (score_white > score_black) {
        printf ("White\n");
    } else if (score_white < score_black) {
        printf ("Black\n");
    } else {
        printf ("Draw\n");
    }

    return 0;
}

