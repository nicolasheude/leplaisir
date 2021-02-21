import * as React from 'react';

export default function Form() {
  const [state, setState] = React.useState({
    child_email: "",
    child_prenom: "",
    child_nom: "",
    child_anniversaire: "",
    child_nationalite: "",
    child_sexe: "",
    child_niveau: "",
    child_fft: "",
    child_locationR: "",
    child_repas: "",
    child_formule: "",
    child_semaine: "",

    parent_nom: "",
    parent_prenom: "",
    parent_sexe: "",
    parent_adresse: "",
    parent_ville: "",
    parent_cp: "",
    parent_email: "",
    parent_tel_1: "",
    parent_tel_2: ""
  })

  function handleChange(evt) {
    const value = evt.target.value;
    setState({
      ...state,
      [evt.target.name]: value
    });
  }

  const handleSubmit = (event) => {
      let child_s = {
        email: event.target.child_email.value,
        nom: event.target.child_nom.value,
        prenom: event.target.child_prenom.value,
        anniversaire: event.target.child_anniversaire.value,
        nationalite: event.target.child_nationalite.value,
        sexe: event.target.child_sexe.value,
        niveau: event.target.child_niveau.value,
        fft: event.target.child_fft.value,
        locationR: event.target.child_locationR.value,
        repas: event.target.child_repas.value,
        formule: event.target.child_formule.value,
        semaine: event.target.child_semaine.value
      }

      let parents_s = {
          nom: event.target.parent_nom.value,
          prenom: event.target.parent_prenom.value,
          sexe: event.target.parent_sexe.value,
          adresse: event.target.parent_adresse.value,
          ville: event.target.parent_ville.value,
          cp: event.target.parent_cp.value,
          email: event.target.parent_email.value,
          tel_1: event.target.parent_tel_1.value,
          tel_2: event.target.parent_tel_2.value,
      }

      let body = {
          child: child_s,
          parents: parents_s
      }

      console.log(body)

      fetch('http://0.0.0.0:8080/form', {
          method: 'POST',
          body: JSON.stringify(body)
        }).then(function(response) {
          console.log(response)
          return response;
        }).catch((err)=>console.error(err)) ;
      event.preventDefault();
  }

  return (
    <React.Fragment>

      {/* Partie Gosse */}

      <section className="hero is-link">
        <div className="hero-body">
          <p className="title">
            Votre enfant
          </p>
          <p className="subtitle">
            lorem ipsum
          </p>
        </div>
      </section>

      <form className="is-centered" onSubmit={handleSubmit}>

      <section className="section">
        <div className="columns">

          <div className="column">
            <div className="field">
              <label className="label">Son nom</label>
              <div className="field-body">
                <div className="control">
                  <input type="text" name="child_nom" value={state.child_nom} onChange={handleChange} required/>
                </div>
              </div>
            </div>
          </div>

          <div className="column">
            <div className="field">
              <label className="label">Son prénom</label>
              <div className="field-body">
                <div className="control">
                  <input type="text" name="child_prenom" value={state.child_prenom} onChange={handleChange} required/>
                </div>
              </div>
            </div>
          </div>

          <div className="column">
            <div className="field">
              <label className="label">Son email</label>
              <div className="field-body">
                <div className="control">
                  <input type="email" name="child_email" value={state.child_email} onChange={handleChange} required/>
                </div>
              </div>
            </div>
          </div>

          <div className="column">
            <div className="field">
              <label className="label">Sa nationalité</label>
              <div className="field-body">
                <div className="control">
                  <input type="text" name="child_nationalite" value={state.child_nationalite} onChange={handleChange} required/>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div className="columns">
          <div className="column">
            <div className="field">
              <label className="label">Sa date d'anniversaire</label>
              <div className="control">
                <input type="date" name="child_anniversaire" value={state.child_anniversaire} onChange={handleChange} required/>
              </div>
            </div>
          </div>

          <div className="column">
            <div className="field">
              <label className="label">Son sexe</label>
              <div className="select">
                  <select type="text" name="child_sexe" value={state.child_sexe} onChange={handleChange} required>
                    <option value="homme">Masculin</option>
                    <option value="femme">Féminin</option>
                  </select>
                </div>
            </div>
          </div>

          <div className="column">
            <div className="field">
              <label className="label">Son niveau</label>
              <div className="control">
                <div className="select">
                  <select type="text" name="child_niveau" value={state.child_niveau} onChange={handleChange} required>
                    <option value="debutant">Débutant</option>
                    <option value="intermediaire">Intermédiaire</option>
                    <option value="experimente">Expérimenté</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <div className="column">
            <div className="field">
              <label className="label">FFT</label>
              <div className="select">
                  <select type="text" name="child_fft" value={state.child_fft} onChange={handleChange} required>
                    <option value="oui">Oui</option>
                    <option value="non">Non</option>
                  </select>
                </div>
            </div>
          </div>
        </div>
        </section>

        <section className="hero is-link">
        <div className="hero-body">
          <p className="title">
            Votre projet
          </p>
          <p className="subtitle">
            lorem ipsum
          </p>
        </div>
      </section>

      <section className="section">
        <div className="columns">
          <div className="column">
            <div className="field">
              <label className="label">Location Raquette ?</label>
              <div className="select">
                  <select type="text" name="child_locationR" value={state.child_locationR} onChange={handleChange} required>
                    <option value="oui">Oui</option>
                    <option value="non">Non</option>
                  </select>
                </div>
            </div>
          </div>

          <div className="column">
          <div className="field">
          <label className="label">Repas ?</label>
          <div className="select">
                  <select type="text" name="child_repas" value={state.child_repas} onChange={handleChange} required>
                    <option value="restaurant">Restaurant</option>
                    <option value="pique_nique">Pique-Nique</option>
                    <option value="autre">Autre</option>
                  </select>
                </div>
          </div>
          </div>

          <div className="column">
          <div className="field">
            <label className="label">Formule (Atelier choisi)</label>
            <div className="select">
                  <select type="text" name="child_formule" value={state.child_formule} onChange={handleChange} required>
                    <option value="Mini-Tennis">Mini-Tennis</option>
                    <option value="Tennis Times">Tennis Times</option>
                    <option value="Tennis + Activitée">Tennis + Activitée</option>
                    <option value="Tennis + Wakeboard">Tennis + Wakeboard</option>
                    <option value="Tennis + Voile">Tennis + Voile</option>
                  </select>
                </div>
          </div>
          </div>

          <div className="column">
          <div className="field">
            <label className="label">Semaine</label>
            <div className="select">
                  <select type="text" name="child_semaine" value={state.child_semaine} onChange={handleChange} required>
                    <option value="SemaineA">SemaineA</option>
                    <option value="SemaineB">SemaineB</option>
                    <option value="SemaineC">SemaineC</option>
                    <option value="SemaineD">SemaineD</option>
                    <option value="SemaineE">SemaineE</option>
                    <option value="SemaineF">SemaineF</option>
                    <option value="SemaineG">SemaineG</option>
                    <option value="SemaineH">SemaineH</option>
                  </select>
                </div>
          </div>
          </div>

        </div>
      </section>

        {/* Partie Parents */}

      <section className="hero is-link">
        <div className="hero-body">
          <p className="title">
            Vous
          </p>
          <p className="subtitle">
            lorem ipsum
          </p>
        </div>
      </section>

      <section className="section">
        <div className="columns is-multiline">
        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre nom</label>
          <div className="control">
            <input type="text" name="parent_nom" value={state.parent_nom} onChange={handleChange} required/>
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre prénom</label>
          <div className="control">
            <input type="text" name="parent_prenom" value={state.parent_prenom} onChange={handleChange} required/>
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Père/Mère</label>
          <div className="select">
                  <select type="text" name="parent_sexe" value={state.parent_sexe} onChange={handleChange} required>
                    <option value="homme">Père</option>
                    <option value="femme">Mère</option>
                  </select>
                </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre adresse</label>
          <div className="control">
            <input type="text" name="parent_adresse" value={state.parent_adresse} onChange={handleChange} required/>
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre ville</label>
          <div className="control">
            <input type="text" name="parent_ville" value={state.parent_ville} onChange={handleChange} required/>
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre Code Postal</label>
          <div className="control">
            <input type="text" name="parent_cp" value={state.parent_cp} onChange={handleChange} required/>
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre Email</label>
          <div className="control">
            <input type="email" name="parent_email" value={state.parent_email} onChange={handleChange} required/>
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre Téléphone mobile</label>
          <div className="control">
            <input type="text" name="parent_tel_1" value={state.parent_tel_1} onChange={handleChange} required/>
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
          <label className="label">Votre Téléphone fixe</label>
          <div className="control">
            <input type="text" name="parent_tel_2" value={state.parent_tel_2} onChange={handleChange} />
          </div>
        </div>
        </div>

        <div className="column is-one-quarter">
        <div className="field">
        </div>

        </div>
        </div>
        <input className="button is-primary" type="submit" value="Envoyer" required/>
        </section>
      </form>
    </React.Fragment>
  );
}
