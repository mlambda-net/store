import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import LocalizedStrings from "react-localization";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";

const lang = new LocalizedStrings({
  en: {
    tags:"Tags"
  },
  es: {
    tags: "Etiquetas"
  },
});

const styles = (theme) => ({

})

class Tags extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    lang.setLanguage(this.props.lang);
    return (
      <div></div>
    )
  }
}

export default withUtils(styles)(withTheme(Tags));
