import React from "react";
import {connect} from "react-redux";
import { withTheme} from "@material-ui/core/styles";
import {FormControl, InputLabel, MenuItem, Select} from "@material-ui/core";
import LocalizedStrings from "react-localization";
import CurrencyService from "../services/currency";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";


const lang = new LocalizedStrings({
  en: {
    "USD": "Dollar",
    "CRC": "CR Colon",
    "Currency": "Currency"
  },
  es: {
    "USD": "Dolares",
    "CRC": "Colones",
    "Currency": "Moneda"
  },
});

const styles = (theme) => ({})

class CurrencySelect extends React.Component {

  constructor(props) {
    super(props);
     this.state = {currency:{}}
    this.service = new CurrencyService()

  }

  async componentDidMount() {
    this.service.fetch()
  }

  name(symbol) {
    if (symbol === "USD") {
      return lang.USD
    } else {
      return lang.CRC
    }
  }

  render() {
    lang.setLanguage(this.props.lang);
    return (
      <FormControl fullWidth variant="standard" color="secondary">
        <InputLabel htmlFor="currency">{lang.Currency}</InputLabel>
        <Select id="currency" value={this.props.value} fullWidth onChange={this.props.onCurrency}>
          {this.props.items.map(c => (<MenuItem key={c.Symbol} value={c.Symbol}  >{this.name(c.Symbol)}</MenuItem>))}
        </Select>
      </FormControl>
    )
  }
}

export default connect(s => s.currency)(withUtils(styles)(withTheme(CurrencySelect)));
