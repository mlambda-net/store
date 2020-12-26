import React from "react";
import {connect} from "react-redux";
import { withTheme} from "@material-ui/core/styles";
import LocalizedStrings from "react-localization";
import {FormControl, InputLabel, MenuItem, Select} from "@material-ui/core";
import CategoryService from "../services/category";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";

const lang = new LocalizedStrings({
  en: {
    category: "Category"
  },
  es: {
    category: "Categoria"
  },
});

const styles = (theme) => ({})

class Category extends React.Component {

  constructor(props) {
    super(props);
    this.service = new CategoryService();
  }

  componentDidMount() {
    this.service.fetch()
  }

  render() {
    const {items, value, onChange} = this.props
    lang.setLanguage(this.props.lang);
    return (
      <FormControl fullWidth>
        <InputLabel id="category">{lang.category}</InputLabel>
        <Select labelId="category" value={value} onChange={onChange}>
          <MenuItem key={-1} value=""/>
          {items.map( (c,i) => <MenuItem key={i} value={c.Name}> {c.Name}</MenuItem>)}
        </Select>
      </FormControl>
    )
  }
}

export default connect(s => s.category)(withUtils(styles)(withTheme(Category)));
