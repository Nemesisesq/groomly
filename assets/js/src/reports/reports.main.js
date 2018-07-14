import React, {Component} from 'react'
import List from '@material-ui/core/List'
import ListItem from "@material-ui/core/ListItem";
import axios from 'axios'
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import AddIcon from "@material-ui/icons/Add"
import SaveIcon from "@material-ui/icons/Save"
import EditIcon from "@material-ui/icons/Edit"
import TextField from "@material-ui/core/TextField";
import _ from 'lodash'

const hostUri = "http://localhost:3000/api";

const ReportList = props => {
    const {select, project_reports} = props

    return (
        <List>
            {project_reports.map(item => {
                return (
                    <ListItem button onClick={_ => select(item)} key={item.id}>
                        {item.name}
                        |
                        {item.summary}
                    </ListItem>
                )
            })}
        </List>
    )
}

const ReportDetail = props => {
    const {detail} = props;
    return (
        <List>
            {Object.keys(detail)
                .filter(x => !_.includes(['id', 'created_at', 'updated_at'], x))
                .map(key => {
                    return (
                        <ListItem key={key}>
                            <TextField
                                id={key}
                                label={key}
                                value={detail[key]}
                                onChange={data => props.handleChange(key, data)}
                                margin="normal"
                                disabled={!props.editing}
                            />
                        </ListItem>
                    )
                })}
        </List>
    )
}


class ProjectReports extends Component {

    constructor(props) {
        super(props);
        this.state = {
            project_reports: [{name: 'hello world', id: 1}, {name: "wornderful", id: 2}],
            detail: {},
            editing: false,
            updating: false
        }
    }

    componentDidMount() {


        axios({
            method: "get",
            url: `${hostUri}/project_reports`,
            responseType: "application/json",
        })
            .then(data => {

                this.setState({
                    project_reports: [...data.data]
                })
            })
            .catch(error => {
                console.log(error)
            })

    }

    _newReport = () => {
        this.setState({
            editing: true,
            updating: false
        })
        axios({
            method: "get",
            url: `${hostUri}/project_reports/new`,
            responseType: "application/json",
        })
            .then(data => {
                this.setState({
                    detail: {...data.data}
                })
            })
            .catch(error => {
                console.log(error)
            })
    }

    _saveDetail = () => {
// better handling is needed
        if (!this.state.editing) {
            this.setState({
                editing: true
            });
            return
        }

        if (this.state.updating) {
            axios({
                method: "put",
                url: `${hostUri}/project_reports/${this.state.detail.id}`,
                responseType: "application/json",
                data: this.state.detail
            })
                .then(data => {
                    this.setState({
                        editing: false
                    })
                })
                .catch(error => {
                    console.log(error)
                })
        } else {

            axios({
                method: "post",
                url: `${hostUri}/project_reports`,
                responseType: "application/json",
                data: this.state.detail
            })
                .then(data => {
                    this.setState({
                        editing: false
                    })
                })
                .catch(error => {
                    console.log(error)
                })

        }
    }
    _handleChange = (key, data) => {

        let real_data = data.target.value
        switch (key) {
            case 'min_hurdle':
                real_data = parseInt(real_data) || 0;
                break
            case 'max_hurdle':
                real_data = parseInt(real_data) || 0;
                break;
            case 'rank':
                real_data = parseInt(real_data) || 0;
        }

        this.setState({
            detail: {
                ...this.state.detail,
                ... {[key]: real_data}
            }
        })
        console.log(data)
    }

    _selectDetail = item => {
        this.setState({
            detail: item,
            updating: true
        })
    }

    _cancelEdit = _ => {
        this.setState({
            editing: false
        })
    }

    render() {
        return (<div>
                <Grid container spacing={24}>
                    <Grid item xs={6}>
                        <ReportList project_reports={this.state.project_reports} select={this._selectDetail}/>
                        <Button variant="fab" color="primary" aria-label="add" onClick={this._newReport}>
                            <AddIcon/>
                        </Button>
                    </Grid>
                    <Grid item xs={6}>
                        <ReportDetail detail={this.state.detail} handleChange={this._handleChange}
                                           editing={this.state.editing}/>
                        <Button variant="fab" color="default" aria-label="add" onClick={this._saveDetail}>
                            {this.state.editing ? <SaveIcon/> : <EditIcon/>}
                        </Button>

                        {this.state.editing && (
                            <Button color="secondary" onClick={this._cancelEdit}>
                                cancel
                            </Button>
                        )}

                    </Grid>
                </Grid>
            </div>
        )
    }
}

const styles = theme => ({
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        marginLeft: theme.spacing.unit,
        marginRight: theme.spacing.unit,
        width: 200,
    },
    menu: {
        width: 200,
    },
});


export default ProjectReports