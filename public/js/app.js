(function() {

  var App = Ember.Application.create();

  App.Router.map(function() {
    this.resource('estudantes')
    this.resource('estudante', { path : ":estudante_id"})
    this.resource('about')
  });

  App.Estudante = DS.Model.extend({
    nome: DS.attr('string'),
    matricula: DS.attr('string'),
    created: DS.attr('date'),
  });

  App.EstudantesRoute = Ember.Route.extend({
    model: function() {
      return this.store.find('estudante');
    }
  });

  App.EstudantesController = Ember.ArrayController.extend({
    actions: {
      create: function() {
        var nome = this.get('nome');
        var matricula = this.get('matricula');

        var estudante = this.store.createRecord('estudante', {
          nome: nome,
          matricula: matricula
        });

        this.set('nome', '');
        this.set('matricula', '');

        estudante.save();
      }      
    }
  });

  App.EstudanteController = Ember.ObjectController.extend({
    isEditing: false,
    actions: {
      edit: function(estudante) {
        this.isEditing = true;
      },
      doneEditing: function() {
        this.isEditing = false;
      },
      destroy: function(estudante) {
        console.log("Removendo estudante: " + estudante.get('nome'));
        estudante.destroyRecord();
      }
    }
  });

})();