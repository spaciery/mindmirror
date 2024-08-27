# Forms

- The purpose of a `<form>` is to allow users to input information and send it.
- The `<form>`‘s `action` attribute determines where the form’s information goes.
- The `<form>`‘s `method` attribute determines how the information is sent and processed.
- To add fields for users to input information we use the `<input>` element and set the `type` attribute to a field of our choosing:
    - Setting `type` to `"text"` creates a single row field for text input.
    - Setting `type` to `"password"` creates a single row field that censors text input.
    - Setting `type` to `"number"` creates a single row field for number input.
    - Setting `type` to `"range"` creates a slider to select from a range of numbers.
    - Setting `type` to `"checkbox"` creates a single checkbox that can be paired with other checkboxes.
    - Setting `type` to `"radio"` creates a radio button that can be paired with other radio buttons.
    - Setting `type` to `"text"` and adding the `list` attribute will pair the `<input>` with a `<datalist>` element if the `list` of `<input>` and the `id` of `<datalist>` are the same.
    - Setting `type` to `"submit"` creates a submit button.
- A `<select>` element is populated with `<option>` elements and renders a dropdown list selection.
- A `<datalist>` element is populated with `<option>` elements and works with an `<input>` to search through choices.
- A `<textarea>` element is a text input field that has a customizable area.
- When a `<form>` is submitted, the `name` of the fields that accept input and the `value` of those fields are sent as `name=value` pairs.

Using the `<form>`element in conjunction with the other elements listed above allows us to create sites that take into consideration the wants and needs of our 
users.

# Client side validations

- Client-side validations happen in the browser before information is sent to a server.
- Adding the `required` attribute to an input related element will validate that the input field has information in it.
- Assigning a value to the `min` attribute of a number input element will validate an acceptable minimum value.
- Assigning a value to the `max` attribute of a number input element will validate an acceptable maximum value.
- Assigning a value to the `minlength` attribute of a text input element will validate an acceptable minimum number of characters.
- Assigning a value to the `maxlength` attribute of a text input element will validate an acceptable maximum number of characters.
- Assigning a regex to `pattern` matches the input to the provided regex.
- If validations on a `<form>` do not pass, the user gets a message explaining why and the `<form>` cannot be submitted.