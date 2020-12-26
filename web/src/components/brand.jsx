import React from "react";
import {connect} from "react-redux";
import { withTheme} from "@material-ui/core/styles";
import LocalizedStrings from "react-localization";
import {Autocomplete} from "@material-ui/lab";
import {TextField} from "@material-ui/core";
import BrandService from "../services/brand";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";

const lang = new LocalizedStrings({
  en: {
    brand: "Brand"
  },
  es: {
    brand: "Marca"
  },
});

const styles = (theme) => ({})

class Brand extends React.Component {

  constructor(props) {
    super(props);
    this.service = new BrandService()
  }

  componentDidMount() {
    this.service.fetch()
  }

  render() {
    lang.setLanguage(this.props.lang);
    return (
      <Autocomplete freeSolo
                    options={this.props.items.map(c => c.Name)}
                    value={this.props.value}
                    onChange={(e, value) => this.props.onChange(value)}
                    renderInput={(params =>
                        <TextField {...params} color="secondary" variant="standard" label={lang.brand}/>
                    )}
      />
    )
  }
}

export default connect(s => s.brand)(withUtils(styles)(withTheme(Brand)));
