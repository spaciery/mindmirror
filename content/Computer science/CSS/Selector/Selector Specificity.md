# Selector Specificity

Specificity is the order by which the browser decides which CSS styles will be displayed. A best practice in CSS is to style elements while using the lowest degree of specificity so that if an element needs a new style, it is easy to override.

IDs are the most specific selector in CSS, followed by classes, and finally, type.

```html
<h1 class='headline'>Breaking News</h1>
```

```css
h1 {
  color: red;
}

.headline {
  color: firebrick;
}
```

In the example code above, the color of the heading would be set to `firebrick`,as the class selector is more specific than the type selector. If an ID attribute (and selector) were added to the code above, the styles within the ID selector’s body would override all other styles for the heading.