# Overflow

All of the components of the [box model](https://www.codecademy.com/resources/docs/css/box-model) comprise an element’s size. For example, an image that has the following dimensions is 364 pixels wide and 244 pixels tall.

- 300 pixels wide
- 200 pixels tall
- 10 pixels [padding](https://www.codecademy.com/resources/docs/css/padding) on the left and right
- 10 pixels padding on the top and bottom
- 2 pixels [border](https://www.codecademy.com/resources/docs/css/borders/border) on the left and right
- 2 pixels border on the top and bottom
- 20 pixels [margin](https://www.codecademy.com/resources/docs/css/margins/margin) on the left and right
- 10 pixels margin on the top and bottom

The total dimensions (364px by 244px) are calculated by adding all of the vertical dimensions together and all of the horizontal dimensions together. Sometimes, these components result 
in an element that is larger than the parent’s containing area. How can we ensure that we can view all of an element that is larger than its parent’s containing area?

The `overflow` property controls what happens to content that spills, or overflows, outside its box. The most commonly used values are:

- `hidden`—when set to this value, any content that overflows will be hidden from view.
- `scroll`—when set to this value, a scrollbar will be added to the element’s box so
that the rest of the content can be viewed by scrolling.
- `visible`—when set to this value, the overflow content will be displayed outside of
the containing element. Note, this is the default value.

```
p {
  overflow: scroll;
}

```

The overflow property is set on a parent element to instruct a web browser on how to render child elements. For example, if a div’s overflow property is set to `scroll`, all children of this div will display overflowing content with a scroll bar.