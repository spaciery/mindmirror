# Attribute selector

The *attribute selector* can be used to target HTML elements that already contain attributes. Elements of the same type can be targeted differently by their attribute or attribute value. This alleviates the need to add new code, like the `class` or `id` attributes.

Attributes can be selected similarly to types, classes, and IDs.

```css
[href]{
   color: magenta;
}
```

The most basic syntax is an attribute surrounded by square brackets. In the above example: `[href]` would target all elements with an `href` attribute and set the [`color`](https://www.codecademy.com/resources/docs/css/colors/color) to `magenta`.And it can get [more granular](https://developer.mozilla.org/en-US/docs/Web/CSS/Attribute_selectors#syntax) from there by adding type and/or attribute values. One way is by using `type[attribute*=value]`.In short, this code selects an element where the attribute contains any instance of the specified value. Let’s take a look at an example.

```html
<img src='/images/seasons/cold/winter.jpg'>
<img src='/images/seasons/warm/summer.jpg'>
```

The HTML code above renders two [`<img>`](https://www.codecademy.com/resources/docs/html/elements/img?page_req=catalog) elements, each containing a `src` attribute with a value equaling a link to an image file.

```css
img[src*='winter'] {
  height: 50px;
}

img[src*='summer'] {
  height: 100px;
}
```

Now take a look at the above CSS code. The *attribute selector* is used to target each image individually.

- The first ruleset looks for an `img` element with an attribute of `src` that contains the string `'winter'`, and sets the [`height`](https://www.codecademy.com/resources/docs/css/sizing/height) to `50px`.
- The second ruleset looks for an `img` element with an attribute of `src` that contains the string `'summer'`, and sets the `height` to `100px`.

Notice how no new HTML markup (like a class or id) needed to be added, and we were still able to modify the styles of each image independently. This is one advantage to using the 
attribute selector!