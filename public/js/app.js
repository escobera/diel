(function() {

  var App = Ember.Application.create();

  App.Router.map(function() {
    this.resource('estudantes', function () {
      this.resource('estudante', { path : ":estudante_id"})
    })
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
      createEstudante: function() {
        var nome = this.get('nome');
        var matricula = this.get('matricula');

        var estudante = this.store.createRecord('estudante', {
          nome: nome,
          matricula: matricula
        });
        this.set('nome', '');
        this.set('matricula', '');

        estudante.save();
      },
      destroyEstudante: function(estudante) {
        console.log("Removendo estudante: " + estudante.get('nome'));
        estudante.destroyRecord();
      }
    }
  });

  App.EstudantesController = Ember.ArrayController.extend({
    actions: {
      edit: function() {
        var nome = this.get('nome');
        var matricula = this.get('matricula');

        var estudante = this.store.createRecord('estudante', {
          nome: nome,
          matricula: matricula
        });
        this.set('nome', '');
        this.set('matricula', '');

        estudante.save();
      },
      destroyEstudante: function(estudante) {
        console.log("Removendo estudante: " + estudante.get('nome'));
        estudante.destroyRecord();
      }
    }
  });

})();