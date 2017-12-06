/*
Package yac helps to analyze icons and give out up to 4 colors (for example for
background, font and button of app card) based and complementary icons colors.

Initial Implementation

The color for the background of the card was selected automatically based on the
icon, the button - translucent white. The algorithm tried to determine the main
color of the icon, sorting the pixels by hue. Such an approach did not always
give a beautiful result, it had shortcomings:

    wrong definition of color,
    "dirty" colors due to averaging,
    dim buttons, boring cards.

What really wanted to

The card was to be a real continuation of the icon. The colors are juicy and
bright. We wanted to create the feeling that the card was carefully done by hand,
and not slipped something carelessly generated automatically.

We want to do more beautifully always, but the resources are not unlimited. To
allocate a command to write a miracle library by definition of colors was not
planned. So, the task:

    Minimal forces to improve the algorithm for determining colors, to figure
    out how to paint a card beautifully, without inventing a spaceship.

New algorithm for determining colors

Step 1

Take the icon. We discard white, black and transparent pixels.

Step 2

Reduce the resulting image to a size of 2 × 2 pixels (with anti-aliasing
disabled). As a result, we get four colors for the icon. If the original picture
is homogeneous, they can be repeated - it's okay.

We have disabled anti-aliasing, so that colors do not mix, do not become "dirty".

In fact, it turns out like this: the square is divided into four parts, we take
the average pixel from the top row of each quarter. In the implementation of
everything is simple: we do not even need a real downsample image and generally
work with graphics. Pixels with the desired position are taken from a
one-dimensional array obtained after the first step.

Step 3

Almost everything is ready. Remained quite a bit: get the resulting colors,
translate to HSL, sort by lightness (L). We are painting a card.

Light scheme:

    background - the lightest color;
    button - closest to the light;
    text - the darkest.

Dark scheme (if two or more colors are dark):

    background - the darkest color;
    button - closest to the dark one;
    text - the lightest.

Applying colors, check the contrast: Lightness difference between background and
button ≥ 20; between the background and the text ≥ 60. If not, correct.

Result

We have got colorful cards, from real colors of icons, without "dirty"
impurities. Due to the use of several colors, the card looks much livelier. It
is especially pleasant that with a homogeneous icon background the card becomes
its direct continuation: the border between them is not noticeable at all.

And most importantly: we provided for special cases:

    Icon from the same color: make the background a little darker so that it does
    not merge.
    Icon with background: look at pixels around the edges; if everyone is the
    same, we put the same background of the card.


Source: https://medium.com/@iammishaanikin/98b8f0dcfdc5
*/
package yac // import "github.com/toby3d/yac"
