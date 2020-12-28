var MessageVisibleTimeout = 5000, // 5s
  SMSCodeDeadline = 2 * 60 * 1000; // 2m

// Добавляем метод в jQuery
(function ($) {
  $.fn.extend({
    message: function (type, text) {
      // Находим контейнер с сообщениями
      let messages = document.querySelector('.messages')
      // Если его нет, создаем 
      if (!messages) {
        messages = document.createElement('div')
        messages.classList.add('messages')
        document.body.appendChild(messages);
      }
      // Создаем сообщение
      let $msg = document.createElement('div'),
        header = type == true ? "Success" : "Error";
      $msg.className = "alert alert-dismissible fade show";
      type == true ? $msg.classList.add("alert-success") : $msg.classList.add("alert-danger")
      $msg.innerHTML = `
            <h4 class="alert-heading">${header}</h4>
            <p>${text}</p>
            <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        `;
      messages.appendChild($msg);
      setTimeout(function () {
        $msg.remove()
      }, MessageVisibleTimeout)
    },
    fail: function (msg) {
      let alert = $(this).find('.alert')
      alert.text(msg)
      alert.removeClass('d-none')
      setTimeout(function () {
        alert.addClass('d-none')
      }, MessageVisibleTimeout)
    },
    auth(responce) {
      // form = this
      if (responce.code > 400) {
        return $(this).fail(responce.data.message)
      }
      // Смс верификация прошла успешно
      if (responce.code == 200 && responce.data.verify) return document.location.href = '/'

      if (responce.code == 200 && responce.data.auth) {
        this.addClass('d-none')
        $('.verify').removeClass('d-none')
        var deadline = new Date(Date.parse(new Date()) + SMSCodeDeadline); // for endless timer
        initializeClock("countdown", deadline);
      }
    }
  })
}(jQuery));

document.addEventListener("DOMContentLoaded", function () {

  // Сохранение данных пользователя
  $("form[data-event='update']").submit(function(e) {
    e.preventDefault();
    if (this.checkValidity() === false) return this.classList.add('was-validated');
    let self = this
    fetch( this.action, {
      method: 'POST',
      headers: { "Content-type": "application/json; charset=UTF-8" },
      body: formToJson($(this).serializeArray())
    })
      .then(
        function(response) {
          if (response.status === 201) {  
            return document.location.href = self.action;
          }
          response.json().then( (res) => $().message(false, res.data))
        }
      )
      .catch(function(err) {
        response.json().then( (res) => $().message(false, err))
      });
  });
  // Удаление пользователя
  $('[data-event="delete"]').on('click', function(e) {
    e.preventDefault();
    let $el = e.target.closest('.item');
    console.log(this.action)
    if ($el && $el.dataset.id) {
      fetch( this.action, {
        method: 'DELETE',
        headers: { "Content-type": "application/json; charset=UTF-8" },
        body: JSON.stringify({ "id": parseInt($el.dataset.id) })
      })
        .then((response) => {
            response.json().then( (res) => $().message(true, res.data))
        })
        .catch(function(err) {
          response.json().then( (res) => $().message(false, res.data))
        });
    }
  });
  // Утвердить карту
  $('[data-event="approve"]').on('click', approveCard);

  $('.page-item').on('click', function (e) {
    let page = e.target.dataset.page
    $.ajax({
        method: 'get',
        url: '/?page=' + page
      })
      .done(function (res) {
        res = JSON.parse(res)
        // console.log(res)
        drawTableClients(res.data)
      })
      .fail(function (e) {
        console.log(e)
      });
  });

  // Форматируем номер телефона
  $('#phone').on('keypress', function (e) {
    IMask(this, {
      mask: '+{7}(000)000-00-00'
    });
  })
  $('#phone').on('blur', function (e) {
    IMask(this, {
      mask: '+{7}(000)000-00-00'
    });
  })

  // Переключатели типа базы данных
  $('input[type=radio]').on('change', function (e) {
    if (this.value == "sqlite3") {
      $('.database_server').addClass('d-none')
      $('.database_path').removeClass('d-none')
    } else {
      $('.database_path').addClass('d-none')
      $('.database_server').removeClass('d-none')
    }
  });



  // Сохранение данных пользователя
  $(".adduser").submit(function (e) {
    e.preventDefault();
    if (this.checkValidity() === false) return this.classList.add('was-validated');
    let self = this
    fetch('/admin/user', {
      method: 'post',
      headers: { "Content-type": "application/json; charset=UTF-8" },
      body: formToJson($(this).serializeArray())
    })
      .then(
        function(response) {
          console.log(response)
        }
      )
      .catch(function(err) {
        console.log('Fetch Error :-S', err);
      });
  });



  // Сохранение настроект
  $(".settings").submit(function (event) {
    event.preventDefault();
    if (this.checkValidity() === false) return this.classList.add('was-validated');

    // Вынести в отдельную функцию | Придумать более красивый вариант
    let data = new FormData(this),
      map = {};
    data.forEach((value, key) => {
      let k = key.split("_")
      if (k.length > 1) {
        if (!map[k[0]]) map[k[0]] = {}
        map[k[0]][k[1]] = value
      } else {
        map[key] = value
      }
    })

    // console.log(map)
    $.post(this.action, JSON.stringify(map))
      .done(responce => $().message(true, responce))
      .fail(error => $().message(false, error));
  });



  // Авторизация и верификация
  $(".login").submit(function (e) {
    let self = this;
    e.preventDefault();
    if (this.checkValidity() === false) return this.classList.add('was-validated');
    $.post(this.action, $(this).serializeArray())
      .done(function (res) {
        let responce = JSON.parse(res)
        $(self).auth(responce)
      })
      .fail(function (e) {
        let error = JSON.parse(e.responseText)
        $(self).fail(error.data.message)
      });
  });

  // Установка приложения
  var navListItems = $('div.setup-panel a'),
    allWells = $('.setup-content'),
    allNextBtn = $('.nextBtn'),
    allBackBtn = $('.backBtn');

  // Скрываем все блоки
  allWells.hide();

  // Обработчики нажатия переключателей шагов
  navListItems.click(function (e) {
    e.preventDefault();
    var $target = $($(this).attr('href')),
      $item = $(this);

    if (!$item.hasClass('disabled')) {
      navListItems.removeClass('btn-primary').addClass('btn-default');
      $item.addClass('btn-primary');
      allWells.hide();
      $target.show();
      $target.find('input:eq(0)').focus();
    }
  });

  // Обработчик кнопк Далее
  allNextBtn.click(function () {
    var curStep = $(this).closest(".setup-content"),
      curStepBtn = curStep.attr("id"),
      nextStepWizard = $('div.setup-panel a[href="#' + curStepBtn + '"]').next(),
      curInputs = curStep.find("input[type='text'],input[type='url']"),
      isValid = true;

    $(".form-control").removeClass("is-invalid");
    for (var i = 0; i < curInputs.length; i++) {
      if (!curInputs[i].validity.valid) {
        isValid = false;
        $(curInputs[i]).closest(".form-control").addClass("is-invalid");
      }
    }

    if (isValid)
      nextStepWizard.removeAttr('disabled').trigger('click');
  });

  // Обработчик кнопк Назад
  allBackBtn.click(function () {
    var curStep = $(this).closest(".setup-content"),
      curStepBtn = curStep.attr("id"),
      nextStepWizard = $('div.setup-panel a[href="#' + curStepBtn + '"]').prev(),
      curInputs = curStep.find("input[type='text'],input[type='url']");

    nextStepWizard.removeAttr('disabled').trigger('click');
  });

  // Показать стартовый блок
  $('div.setup-panel a.btn-primary').trigger('click');

  // Обработчик выбора драйвера БД
  $('.db-drivers').click(function (e) {
    let type = $(e.target).children('input').val();

    if (type == 'sqlite3') {
      $('.bd_server').hide()
      $('.bd_path').show()
    } else {
      $('.bd_server').show()
      $('.bd_path').hide()
    }
  });

  $('[data-event="install"]').submit(function (e) {
    e.preventDefault();

    // Вынести в отдельную функцию
    let data = new FormData(this),
      map = {};
    data.forEach((value, key) => {
      let k = key.split("_")
      if (k.length > 1) {
        if (!map[k[0]]) map[k[0]] = {}
        map[k[0]][k[1]] = value
      } else {
        map[key] = value
      }
    })

    $.post(this.action, JSON.stringify(map))
      .done(responce => $().message(true, responce))
      .fail(error => $().message(false, error));
  });

});

function formToJson(data) {
  let json = {}
  data.forEach(function (el) {
    json[el.name] = el.value
  });
  json.id = parseInt(json.id)
  return JSON.stringify(json)
}

function getTimeRemaining(endtime) {
  var t = Date.parse(endtime) - Date.parse(new Date());
  var seconds = Math.floor((t / 1000) % 60);
  var minutes = Math.floor((t / 1000 / 60) % 60);
  return {
    total: t,
    minutes: minutes,
    seconds: seconds
  };
}

function initializeClock(id, endtime) {
  console.log(endtime)
  var clock = document.getElementById(id);
  var minutesSpan = clock.querySelector(".minutes");
  var secondsSpan = clock.querySelector(".seconds");

  function updateClock() {
    var t = getTimeRemaining(endtime);

    if (t.total <= 0) {
      $('#countdown').addClass("d-none");
      $('#sms-repeat').removeClass("d-none");
      clearInterval(timeinterval);
      return true;
    }

    minutesSpan.innerHTML = ("0" + t.minutes).slice(-2);
    secondsSpan.innerHTML = ("0" + t.seconds).slice(-2);
  }

  updateClock();
  var timeinterval = setInterval(updateClock, 1000);
}

function drawTableClients(clients) {

  document.getElementById('tbody').innerHTML = ''

  drawPagginator(clients.length)

  clients.forEach((client) => {
    tr = document.createElement('tr');

    tr.classList.add('item')
    tr.dataset.clientid = client.ClientID
    tr.dataset.depositid = client.DepositID
    tbody.append(tr)

    td = document.createElement('td')
    td.innerText = client.FullName
    tr.append(td)

    td = document.createElement('td')
    td.innerText = client.Bill
    tr.append(td)

    td = document.createElement('td')
    td.innerText = client.OrgName
    tr.append(td)

    td = document.createElement('td')
    td.innerText = client.PhoneNumber
    tr.append(td)

    td = document.createElement('td')
    td.innerHTML = '<button data-approve="approve" data-event="approve" class="btn btn-success btn-sm">Утвердить</button>'
    td.addEventListener('click', approveCard)
    tr.append(td)

    console.log(client)
  })

}

function drawPagginator(length) {
  document.getElementById('pagination').innerHTML = ''

  for (i = 1; i < length + 1; i++) {
    li = document.createElement('li');
    li.classList.add('page-item')

    li.innerHTML = '<span class="page-link" data-page="' + i + '">' + i + '</span>';

    document.getElementById('pagination').append(li)
  }

}

function approveCard() {
  let $el = $(this).closest('.item');
  let data = {
      "ClientID": $($el).data('clientid'),
      "DepositID": $($el).data('depositid'),
      "PhoneNumber": $($el).find(".PhoneNumber").val()
    },
    self = this;
  $.post('/approve', JSON.stringify(data))
    .done(responce => {
      let res = JSON.parse(responce)
      // clients = res.clients;
      $().message(true, res.status)
      $($el).fadeOut();
    })
    .fail(error => $().message(false, error.status));
}