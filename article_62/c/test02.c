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

    window = SDL_CreateWindow("SDL2 example #2", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, SDL_WINDOW_SHOWN);

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
        SDL_Rect rects[1];
        rects[0].x = WIDTH/4;
        rects[0].y = HEIGHT/4;
        rects[0].w = WIDTH/2;
        rects[0].h = HEIGHT/2;

        SDL_UpdateWindowSurfaceRects(window, rects, 1);
    }

    SDL_Delay(5000);

    SDL_DestroyWindow(window);

    SDL_Quit();

    return 0;
}
