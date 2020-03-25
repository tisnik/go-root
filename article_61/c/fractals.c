#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <SDL2/SDL.h>

#define WIDTH 320
#define HEIGHT 240

#define nil NULL

#define MAX(a,b) ((a)>(b) ? (a) : (b))
#define MIN(a,b) ((a)<(b) ? (a) : (b))

SDL_Surface *pixmap;
SDL_Surface *screen_surface = NULL;
SDL_Surface *bitmap_font_surface = NULL;
SDL_Window *window = NULL;

double xpos = -0.75;
double ypos = 0.0;
double scale = 150.0;
double  uhel = 45.0;

int gfx_initialize(int fullscreen, int width, int height, int bpp) {
    window = NULL;
    if (SDL_Init(SDL_INIT_VIDEO) < 0) {
        fprintf(stderr, "Error initializing SDL: %s\n", SDL_GetError());
        return 1;
    } else {
        puts("SDL_Init ok");
    }

    window = SDL_CreateWindow( "Example", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, width, height, SDL_WINDOW_SHOWN );
    if (!window) {
        puts("Error creating window");
        puts(SDL_GetError());
        return 1;
    } else {
        puts("SDL_CreateWindow ok");
    }

    screen_surface = SDL_GetWindowSurface(window);

    if (!screen_surface) {
        fprintf(stderr, "Error setting video mode: %s\n", SDL_GetError());
        return 1;
    } else {
        puts("SDL_GetWindowSurface ok");
    }
    return 0;
}

void gfx_finalize(void) {
    SDL_FreeSurface(screen_surface);
    SDL_DestroyWindow(window);
}

SDL_Surface* gfx_create_surface(int width, int height) {
    return SDL_CreateRGBSurface(SDL_SWSURFACE, width, height, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0x00000000);
}

void putpixel(SDL_Surface *surface, int x, int y, unsigned char r, unsigned char g, unsigned char b)
{
    if (x>=0 && x< surface->w && y>=0 && y < surface->h) {
        if (surface->format->BitsPerPixel == 24) {
            Uint8 *pixel = (Uint8 *)surface->pixels;
            pixel += x*3;
            pixel += y*surface->pitch;
            *pixel++ = b;
            *pixel++ = g;
            *pixel   = r;
        }
        if (surface->format->BitsPerPixel == 32) {
            Uint8 *pixel = (Uint8 *)surface->pixels;
            pixel += x*4;
            pixel += y*surface->pitch;
            *pixel++ = b;
            *pixel++ = g;
            *pixel   = r;
        }
    }
}

void hline(SDL_Surface *surface, int x1, int x2, int y, unsigned char r, unsigned char g, unsigned char b) {
    int x;
    int fromX = MIN(x1, x2);
    int toX = MAX(x1, x2);
    for (x = fromX; x <= toX; x++) {
        putpixel(surface, x, y, r, g, b);
    }
}

void vline(SDL_Surface *surface, int x,  int y1, int y2, unsigned char r, unsigned char g, unsigned char b) {
    int y;
    int fromY = MIN(y1, y2);
    int toY = MAX(y1, y2);
    for (y = fromY; y <= toY; y++) {
        putpixel(surface, x, y, r, g, b);
    }
}

static void show_fractal(SDL_Surface *surface) {
    SDL_BlitSurface(surface, NULL, screen_surface, NULL);
    SDL_UpdateWindowSurface(window);
}

void calcCorner(double xpos, double ypos, double scale,
                double *xmin,  double *ymin,  double *xmax, double *ymax) {
    *xmin=xpos-WIDTH/scale;
    *ymin=ypos-HEIGHT/scale;
    *xmax=xpos+WIDTH/scale;
    *ymax=ypos+HEIGHT/scale;
}

void draw_grid(SDL_Surface *surface) {
    int width = surface->w;
    int height = surface->h;
    int x, y;
    SDL_FillRect(surface, NULL, 0xffffffff);

    for (x=0; x<width; x+=20) {
        vline(surface, x, 0, height-1, 191, 191, 255);
    }
    for (y=0; y<height; y+=20) {
        hline(surface, 0, width-1, y, 191, 191, 255);
    }
}

void draw_mandeljulia(SDL_Surface *surface) {
    double  zx,zy,zx2,zy2,cx,cy;
    double  cosu,sinu,ccxc,ccyc;
    int     x,y,i;
    Uint8 *pixel = nil;

    double ccx = 0.0;
    double ccy = 0.0;

    double xmin, ymin, xmax, ymax;
    double u;

    calcCorner(xpos, ypos, scale, &xmin, &ymin, &xmax, &ymax);

    u=uhel*3.1415/180.0;
    cosu=cos(u);
    sinu=sin(u);
    ccxc=ccx*cosu;
    ccyc=ccy*cosu;

    cy = ymin;

    for (y=0;y<240;y++) {
        cx=xmin;
        pixel = (Uint8 *)surface->pixels + (y + 128) * surface->pitch + 160*4;
        for (x=0;x<320;x++) {
            i=0;
            zx=cx*cosu;
            zy=cy*cosu;
            do {
                zx2=zx*zx;
                zy2=zy*zy;
                zy=2.0*zx*zy+ccyc+cy*sinu;
                zx=zx2-zy2+ccxc+cx*sinu;
                i++;
            } while (i<64 && zx2+zy2<4.0);
            {
                int r = i*2;
                int g = i*3;
                int b = i*5;

                *pixel++ = r;
                *pixel++ = g;
                *pixel++ = b;
                pixel++;
            }
            cx += (xmax-xmin)/WIDTH;
        }
        cy += (ymax-ymin)/HEIGHT;
    }
}

void redraw(SDL_Surface *pixmap) {
    draw_grid(pixmap);
    draw_mandeljulia(pixmap);
    show_fractal(pixmap);
}

static void main_event_loop(void) {
    SDL_Event event;
    int done = 0;
    int left = 0, right = 0, up = 0, down = 0;
    int zoomin = 0, zoomout = 0;
    int perform_redraw;
    int angle_1 = 0, angle_2 = 0;

    do {
        /*SDL_WaitEvent(&event);*/
        while (SDL_PollEvent(&event)) {
            switch (event.type) {
                case SDL_QUIT:
                    done = 1;
                    break;
                case SDL_KEYDOWN:
                    switch (event.key.keysym.sym) {
                        case SDLK_ESCAPE:
                        case SDLK_q:
                            done = 1;
                            break;
                        case SDLK_LEFT:
                            left = 1;
                            break;
                        case SDLK_RIGHT:
                            right = 1;
                            break;
                        case SDLK_UP:
                            up = 1;
                            break;
                        case SDLK_DOWN:
                            down = 1;
                            break;
                        case SDLK_PAGEDOWN:
                            zoomin = 1;
                            break;
                        case SDLK_PAGEUP:
                            zoomout = 1;
                            break;
                        case SDLK_z:
                            angle_1 = 1;
                            break;
                        case SDLK_x:
                            angle_2 = 1;
                            break;
                        default:
                            break;
                    }
                    break;
                case SDL_KEYUP:
                    switch (event.key.keysym.sym) {
                        case SDLK_LEFT:
                            left = 0;
                            break;
                        case SDLK_RIGHT:
                            right = 0;
                            break;
                        case SDLK_UP:
                            up = 0;
                            break;
                        case SDLK_DOWN:
                            down = 0;
                            break;
                        case SDLK_PAGEDOWN:
                            zoomin = 0;
                            break;
                        case SDLK_PAGEUP:
                            zoomout = 0;
                            break;
                        case SDLK_z:
                            angle_1 = 0;
                            break;
                        case SDLK_x:
                            angle_2 = 0;
                            break;
                        default:
                            break;
                    }
                default:
                    break;
            }
        }
        perform_redraw = 0;
        if (left) {
            xpos -= 10.0/scale;
            perform_redraw=1;
        }
        if (right) {
            xpos += 10.0/scale;
            perform_redraw=1;
        }
        if (up) {
            ypos -= 10.0/scale;
            perform_redraw=1;
        }
        if (down) {
            ypos += 10.0/scale;
            perform_redraw=1;
        }
        if (zoomin) {
            scale *= 0.9;
            perform_redraw=1;
        }
        if (zoomout) {
            scale *= 1.1;
            perform_redraw=1;
        }
        if (angle_1) {
            uhel--;
            perform_redraw=1;
        }
        if (angle_2) {
            uhel++;
            perform_redraw=1;
        }
        if (perform_redraw) {
            redraw(pixmap);
        }
    } while (!done);
}

int main(int argc, char **argv) {
    if (gfx_initialize(0, 640, 480, 32)) {
        return 1;
    }

    pixmap = gfx_create_surface(screen_surface->w, screen_surface->h);

    draw_grid(pixmap);
    draw_mandeljulia(pixmap);
    show_fractal(pixmap);
    main_event_loop();

    gfx_finalize();
    SDL_Quit();
    return 0;
}
