#include <stdio.h>
#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

#define WIDTH  640
#define HEIGHT 480

int main(int argc, char** args) {
    SDL_Surface *primarySurface = NULL;
    SDL_Window  *window = NULL;
    SDL_Surface *image;

    if (SDL_Init(SDL_INIT_VIDEO) < 0) {
        puts("Error initializing SDL");
        puts(SDL_GetError());
        return 1;
    } 

    window = SDL_CreateWindow("SDL2 example #13", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, SDL_WINDOW_SHOWN);

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

    image = IMG_Load("globe.png");

    if (!image) {
        puts("Error loading image");
        return 0;
    }

    {
        SDL_Rect dstRect;
        dstRect.y = 10;
        int y;

        for (y=0; y<4; y++) {
            SDL_SetSurfaceAlphaMod(image, y*64);

            SDL_SetSurfaceBlendMode(image, SDL_BLENDMODE_NONE);
            dstRect.x = 10;
            SDL_BlitSurface(image, NULL, primarySurface, &dstRect);

            SDL_SetSurfaceBlendMode(image, SDL_BLENDMODE_BLEND);
            dstRect.x += 100;
            SDL_BlitSurface(image, NULL, primarySurface, &dstRect);

            SDL_SetSurfaceBlendMode(image, SDL_BLENDMODE_ADD);
            dstRect.x += 100;
            SDL_BlitSurface(image, NULL, primarySurface, &dstRect);

            SDL_SetSurfaceBlendMode(image, SDL_BLENDMODE_MOD);
            dstRect.x += 100;
            SDL_BlitSurface(image, NULL, primarySurface, &dstRect);

            dstRect.y += 100;
        }
    }

    SDL_UpdateWindowSurface(window);

    SDL_Delay(5000);

    SDL_FreeSurface(image);
    SDL_DestroyWindow(window);
    SDL_Quit();

    return 0;
}
