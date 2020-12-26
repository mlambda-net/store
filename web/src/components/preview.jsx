import React from "react";
import {withTheme} from "@material-ui/core/styles";
import LocalizedStrings from "react-localization";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";
import Box from "@material-ui/core/Box";
import {IconButton, Typography} from "@material-ui/core";
import { PhotoCamera} from "@material-ui/icons";
import ImageService from "../services/image";

const lang = new LocalizedStrings({
  en: {
    preview: "Preview"
  },
  es: {
    preview: "Miniatura"
  },
});

const styles = (theme) => ({
  thumb: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  preview: {
    display: 'flex',
    flexDirection: 'column',
    width: '100%'
  },
  previewImage: {
    display: 'flex',
    justifyContent: 'flex-start',
    padding: theme.spacing(1,0),
  },
  previewActions: {
    display: 'flex',
    alignItems: 'center'
  },
  imageItem : {
    margin: theme.spacing(1)
  },
  uploadIcon: {
    padding: theme.spacing(0,2)
  },
  imageView : {
    position: 'relative',
    width: "200px",
    height: "160px"
  },
  previewClose: {
    position: 'absolute',
    right: '0',
  },

})

class Preview extends React.Component {
  constructor(props) {
    super(props);
    this.service = new ImageService()
    this.fileSelect = this.fileSelect.bind(this)
  }



  fileSelect(e) {
    const file = e.target.files[0]
    let data = new FormData();
    data.append('file', file)
    data.append('id', this.props.id)
    this.service.save(data).then(c => {
      if(this.props.onChange !== null) {
        this.props.onChange(c)
      }
    })
  }

  render() {
    lang.setLanguage(this.props.lang);
    const {classes} = this.props
    return (
      <div className={classes.thumb}>
        <div className={classes.preview}>
          <Box className={classes.previewActions}>
            <Typography variant="subtitle1" color="textSecondary">
              {lang.preview}
            </Typography>
            <label htmlFor="preview" className={classes.uploadIcon}>
              <IconButton color="secondary" aria-label="upload picture" component="span" size="small">
                <PhotoCamera/>
              </IconButton>
            </label>
          </Box>
          <div className={classes.previewImage}>
            <div className={classes.imageView}>
              <input accept="image/*" type="file" onChange={this.fileSelect} id="preview" style={{display: 'none',}}/>
              <img src={this.props.value} className={classes.imageItem} width="100%" height="100%" alt="preview"/>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default withUtils(styles)(withTheme(Preview));
