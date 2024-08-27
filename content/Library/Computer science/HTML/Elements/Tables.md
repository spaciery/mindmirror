# Tables

A very common task in HTML is structuring tabular data, and it has a number of elements and attributes for just this purpose.

A table is a structured set of data made up of rows and columns (**tabular data**). A table allows you to quickly and easily look up values that indicate some kind of connection between different types of data, for example a person and their age, or a day of the week, or the timetable for a local swimming pool.

- The `<table>` element creates a table.
- The `<tr>` element adds rows to a table.
- To add data to a row, you can use the `<td>` element.
- Table headings clarify the meaning of data. Headings are added with the `<th>` element.
- Table data can span columns using the `colspan` attribute.
- Table data can span rows using the `rowspan` attribute.
- Tables can be split into three main sections: a head, a body, and a footer.
- A table’s head is created with the `<thead>` element.
- A table’s body is created with the `<tbody>` element.
- A table’s footer is created with the `<tfoot>` element.
- All the CSS properties you learned about in this course can be applied to tables and their data.

```html
<table>
  <tbody>
    <tr>
      <th></th>
      <th>Saturday</th>
      <th>Sunday</th>
    </tr>
    <tr>
      <th>Morning</th>
      <td rowspan="2">Work</td>
      <td rowspan="3">Relax</td>
    </tr>
    <tr>
     <th>Afternoon</th>
    </tr>
    <tr>
      <th>Evening</th>
      <td>Dinner</td>
    </tr>
  </tbody>
</table>
```