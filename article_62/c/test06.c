#include <stdio.h>
#include <SDL2/SDL.h>

#define WIDTH  640
#define HEIGHT 480

int main(int argc, char** args) {
    SDL_Surface *primarySurface = NULL;
    SDL_Window  *window = NULL;
    SDL_Surface *image;
    SDL_Surface *tempImage;

    if (SDL_Init(SDL_INIT_VIDEO) < 0) {
        puts("Error initializing SDL");
        puts(SDL_GetError());
        return 1;
    } 

    window = SDL_CreateWindow("SDL2 example #6", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, SDL_WINDOW_SHOWN);

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

    tempImage = SDL_LoadBMP("test1.bmp");

    if (!tempImage) {
        puts("Error loading image");
        return 0;
    }

    image = SDL_ConvertSurface(tempImage, primarySurface->format, 0);

    SDL_FreeSurface(tempImage);
    {
        SDL_Rect srcRect;
        srcRect.x = image->w/4;
        srcRect.y = 0;
        srcRect.w = image->w/2;
        srcRect.h = image->h/2;

        SDL_Rect dstRect;
        dstRect.x = WIDTH/2;
        dstRect.y = HEIGHT/2;
        dstRect.w = 100;
        dstRect.h = 100;
        SDL_BlitSurface(image, &srcRect, primarySurface, &dstRect);
    }
    SDL_Delay(1000);

    SDL_UpdateWindowSurface(window);

    SDL_Delay(5000);

    SDL_FreeSurface(image);
    SDL_DestroyWindow(window);
    SDL_Quit();

    return 0;
}
