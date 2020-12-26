import React from "react";
import {withTheme} from "@material-ui/core/styles";
import LocalizedStrings from "react-localization";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";
import {IconButton, Typography} from "@material-ui/core";
import {PhotoCamera} from "@material-ui/icons";
import ImageService from "../services/image";

const lang = new LocalizedStrings({
  en: {
    images: "Detail Images"
  },
  es: {
    images: "Images"
  },
});

const styles = (theme) => ({

  images: {
    display: 'flex',
    flexDirection: 'column'
  },

  imagesActions :{
    display: 'flex',
    alignItems:'center'
  },

  imagesItems: {
    padding: theme.spacing(1,0),
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'flex-start',
  },
  imageItem : {
    margin: theme.spacing(1)
  },
  uploadIcon: {
    padding: theme.spacing(0,2)
  }

})

class Images extends React.Component {

  constructor(props) {
    super(props);
    this.selectImages = this.selectImages.bind(this)
    this.service = new ImageService()
  }

  selectImages(e) {
    let images = []
    for (let i = 0; i < e.target.files.length && i < 4; i++) {
      const file = e.target.files[i]
      let data = new FormData();
      data.append('file', file)
      data.append('id', this.props.id)
      this.service.save(data).then(url => {
        images[i] = url
        if (this.props.onChange !== null) {
          this.props.onChange(images)
        }
      })
    }

  }


  render() {
    lang.setLanguage(this.props.lang);
    const {classes} = this.props
    return (
      <div className={classes.images}>
        <div className={classes.imagesActions}>
          <Typography variant="subtitle1" color="textSecondary">
            {lang.images}
          </Typography>
          <label htmlFor="img0" className={classes.uploadIcon}>
            <IconButton color="secondary" aria-label="upload picture" component="span" size="small">
              <PhotoCamera/>
            </IconButton>
          </label>
          <input id="img0" accept="image/*" multiple type="file" onChange={this.selectImages} style={{display: 'none'}}/>
        </div>
        <div className={classes.imagesItems}>
          {this.props.value.map((i, k) => <img key={k} src={i} className={classes.imageItem} width="200px" height="160px" alt="catalog"/>)}
        </div>
      </div>
    )
  }
}

export default withUtils(styles)(withTheme(Images));
