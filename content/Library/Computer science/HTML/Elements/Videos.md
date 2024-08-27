# Videos

Unlike the `<img>` element however, the `<video>` element requires an opening and a closing tag.

```html
<video src="myVideo.mp4" width="320" height="240" controls>
  Video not supported
</video>
```

In this example, the video source (`src`) is `"myVideo.mp4"`. The source can be a video file that is hosted alongside your webpage, or a URL that points to a video file hosted on another webpage.

After the `src` attribute, the `width` and `height` attributes are used to set the size of the video displayed in the browser. The `controls` attribute instructs the browser to include basic video controls such as pausing and playing.

The text, `Video not supported`, between the opening and closing video tags will only be displayed if the browser is unable to load the video.