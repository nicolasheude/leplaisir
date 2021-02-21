import React from 'react';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';

export default function TextfieldsCard() {
    return (
        <Grid container spacing={3} style={{
            margin: 10,
            minWidth: 120
            }}>
            <Grid item xs={12} sm={6}>
                <TextField
                    required
                    id="idcarte"
                    name="idcarte"
                    label="Nom et prénom du titulaire de la carte"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
        </Grid>
    );
}