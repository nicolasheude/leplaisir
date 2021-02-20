import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import * as React from 'react';
import { Component } from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        marginLeft: theme.spacing(10),
        marginRight: theme.spacing(10),
        width: 200,
    },
}));

export default function SimpleSelectSemaine() {
    const classes = useStyles();
    const [age, setAge] = React.useState('');

    const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAge(event.target.value as string);
    };

    return (
        <div>
            <FormControl className='select' 
            style={{
                margin: 25,
                minWidth: 120
                }}>
                <InputLabel id="demo-simple-select-label">Semaine</InputLabel>
                <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={age}
                    onChange={handleChange}
                >
                    <MenuItem value={10}>Semaine 1</MenuItem>
                    <MenuItem value={20}>Semaine 2</MenuItem>
                    <MenuItem value={30}>Semaine 3</MenuItem>
                </Select>
            </FormControl>
        </div>
    );
}