# Images

The `<img>` tag allows you to add an image to a web page. Most elements require both opening and closing tags, but the `<img>` tag is a *self-closing tag*. Note that the end of the `<img>` tag has a forward slash `/`. Self-closing tags may include or omit the final slash — both will render properly.

```html
<img src="image-location.jpg" />
```

The `<img>` tag has a required *attribute* called `src`. The `src` attribute must be set to the image’s *source*, or the location of the image. In this case, the value of `src` must be the *uniform resource locator* (URL) of the image. A URL is the web address or local address where a file is stored

The `alt` attribute, which means alternative text, brings meaning to the [images](https://www.codecademy.com/resources/docs/html/images) on our sites. The `alt` attribute can be added to the image tag just like the [`src`](https://www.codecademy.com/resources/docs/html/attributes/src) attribute.  The value of `alt` should be a description of the image.

```html
<img src="#" alt="A field of yellow sunflowers" />
```

The `alt` attribute also serves the following purposes:

- If an image fails to load on a web page, a
user can mouse over the area originally intended for the image and read a brief description of the image. This is made possible by the
description you provide in the `alt` attribute.
- Visually impaired users often browse the web with the aid of screen reading software. When you include the `alt` attribute, the screen reading software can read the image’s description out loud to the visually impaired user.
- The `alt` attribute also plays a role in Search Engine Optimization (SEO),
because search engines cannot “see” the images on websites as they crawl the internet. Having descriptive `alt` [attributes](https://www.codecademy.com/resources/docs/html/attributes) can improve the ranking of your site.

If the image on the web page is not one that conveys any meaningful information to a user (visually impaired or otherwise), the `alt` attribute should be left empty.