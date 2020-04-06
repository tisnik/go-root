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

    window = SDL_CreateWindow("SDL2 example #14", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, SDL_WINDOW_SHOWN);

    if (!window) {
        puts("Error creating window");
        puts(SDL_GetError());
        return 1;
    }

    primarySurface = SDL_GetWindowSurface(window);
    printf("Must lock: %d\n", SDL_MUSTLOCK(primarySurface));

    if (!primarySurface) {
        puts("Error getting surface");
        puts(SDL_GetError());
        return 1;
    }

    {
        int y;

        SDL_LockSurface(primarySurface);
        for (y=0; y<primarySurface->h; y++) {
            int scanLine = primarySurface->pitch;
            Uint8 *ptr = primarySurface->pixels + y*scanLine;
            SDL_memset(ptr, y, scanLine);
        }
        SDL_UnlockSurface(primarySurface);
    }
    SDL_Delay(1000);

    SDL_UpdateWindowSurface(window);

    SDL_Delay(5000);

    SDL_DestroyWindow(window);

    SDL_Quit();
    
    return 0;
}
