import React from 'react';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';

function TextFields() {
    return (
        <Grid container spacing={3} style={{
            margin: 10,
            minWidth: 120
            }}>
            <Grid item xs={12} sm={6}>
                <TextField
                    required
                    id="prenom"
                    name="prenom"
                    label="Prénom de l'enfant"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                    required
                    id="nom"
                    name="nom"
                    label="Nom de l'enfant"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
            <Grid item xs={12}>
                <TextField
                    required
                    id="email"
                    name="email"
                    label="Email"
                    fullWidth
                    autoComplete="off"
                />
            </Grid>
        </Grid>
    );
}