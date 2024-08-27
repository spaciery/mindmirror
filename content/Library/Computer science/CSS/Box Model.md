# Box Model

All elements on a web page are interpreted by the browser as “living” inside of a box. This is what is meant by the box model.

[https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmedia.gcflearnfree.org%2Fcontent%2F5ef2084faaf0ac46dc9c10be_06_23_2020%2Fbox_model.png&f=1&nofb=1&ipt=a5dda56da9760d7b14f494d7d5a9b1612936c2f3df260621ce0030daa7cac580&ipo=images](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmedia.gcflearnfree.org%2Fcontent%2F5ef2084faaf0ac46dc9c10be_06_23_2020%2Fbox_model.png&f=1&nofb=1&ipt=a5dda56da9760d7b14f494d7d5a9b1612936c2f3df260621ce0030daa7cac580&ipo=images)

1. [`width`](https://www.codecademy.com/resources/docs/css/sizing/width) and [`height`](https://www.codecademy.com/resources/docs/css/sizing/height): The width and height of the content area.
2. `padding`: The amount of space between the content area and the border.
3. `border`: The thickness and style of the border surrounding the content area and padding.
4. `margin`: The amount of space between the border and the outside edge of the element.

[**Overflow**](Box%20Model%20c710a02f66564e829dc0b1e5f1f5c4c9/Overflow%200ab77691a5c849d5860796f13a6d7901.md)

1. In the default [box model](https://www.codecademy.com/resources/docs/css/box-model), box dimensions are affected by border thickness and padding.
2. The `box-sizing` property controls the box model used by the browser.
3. The default value of the `box-sizing` property is `content-box`.
4. The value for the new box model is `border-box`.
5. The `border-box` model is not affected by border thickness or padding.