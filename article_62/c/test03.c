#include <stdio.h>
#include <SDL2/SDL.h>

#define WIDTH  640
#define HEIGHT 480

int main(int argc, char** args) {
    SDL_Surface *primarySurface = NULL;
    SDL_Window  *window = NULL;

    if (SDL_Init(SDL_INIT_VIDEO) < 0) {
        puts("Error initializing SDL");
        puts(SDL_GetError());
        return 1;
    } 

    window = SDL_CreateWindow("SDL2 example #3", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, SDL_WINDOW_SHOWN);

    if (!window) {
        puts("Error creating window");
        puts(SDL_GetError());
        return 1;
    }

    primarySurface = SDL_GetWindowSurface(window);

    if (!primarySurface) {
        puts("Error getting surface");
        puts(SDL_GetError());
        return 1;
    }

    SDL_FillRect(primarySurface, NULL, SDL_MapRGB(primarySurface->format, 192, 255, 192));
    SDL_Delay(1000);

    {
#define BORDER 50

        SDL_Rect rects[2];
        rects[0].x = BORDER;
        rects[0].y = BORDER;
        rects[0].w = WIDTH/2;
        rects[0].h = HEIGHT/2;

        rects[1].x = WIDTH-WIDTH/2-BORDER;
        rects[1].y = HEIGHT-HEIGHT/2-BORDER;
        rects[1].w = WIDTH/2;
        rects[1].h = HEIGHT/2;

        SDL_UpdateWindowSurfaceRects(window, rects, 2);
    }

    SDL_Delay(5000);

    SDL_DestroyWindow(window);

    SDL_Quit();

    return 0;
}
