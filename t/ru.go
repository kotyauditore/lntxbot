package t

var RU = map[Key]string{
	NO:         "Нет",
	YES:        "Да",
	CANCEL:     "Отмена",
	CANCELED:   "Отменено.",
	COMPLETED:  "Выполнено!",
	CONFIRM:    "Подтвердить",
	FAILURE:    "Сбой.",
	PROCESSING: "Обрабатываю...",
	WITHDRAW:   "Вывести?",
	ERROR:      "{{if .App}}<b>[{{.App}}]</b> {{end}}Ошибка{{if .Err}}: {{.Err}}{{else}}!{{end}}",
	CHECKING:   "Проверка...",
	TXCANCELED: "Транзакция отменена.",
	UNEXPECTED: "Неожиданная ошибка: пожалуйста, обратитесь в поддержку.",

	CALLBACKWINNER:  "Победитель: {{.Winner}}",
	CALLBACKERROR:   "{{.BotOp}} ошибка{{if .Err}}: {{.Err}}{{else}}.{{end}}",
	CALLBACKEXPIRED: "{{.BotOp}} время истекло.",
	CALLBACKATTEMPT: "Ищу маршрут.",
	CALLBACKSENDING: "Отправляю платёж.",

	INLINEINVOICERESULT:  "Счёт на {{.Sats}} сат.",
	INLINEGIVEAWAYRESULT: "Раздать {{.Sats}}",
	INLINEGIVEFLIPRESULT: "Раздаёт {{.Sats}} сат одному из {{.MaxPlayers}} участников",
	INLINECOINFLIPRESULT: "Лотерея с входным платежом {{.Sats}} сат для {{.MaxPlayers}} участников",
	INLINEHIDDENRESULT:   "{{.HiddenId}} ({{if gt .Message.Crowdfund 1}}crowd:{{.Message.Crowdfund}}{{else if gt .Message.Times 0}}priv:{{.Message.Times}}{{else if .Message.Public}}pub{{else}}priv{{end}}): {{.Message.Content}}",

	LNURLINVALID: "Неверный lnurl: {{.Err}}",
	LNURLFAIL:    "Ошибка при выводе через lnurl: {{.Err}}",

	USERALLOWED:       "Счёт оплачен. {{.User}} допущен.",
	SPAMFILTERMESSAGE: "Привет, {{.User}}. У вас 15мин, чтобы оплатить счёт на {{.Sats}} сат если вы хотите остаться в этой группе:",

	PAYMENTFAILED: "Платёж не состоялся. /log{{.ShortHash}}",
	PAIDMESSAGE: `Оплачено <b>{{.Sats}} сат</b> (+ {{.Fee}} комиссия). 

<b>Hash:</b> {{.Hash}}
{{if .Preimage}}<b>Proof:</b> {{.Preimage}}{{end}}

/tx{{.ShortHash}}`,
	DBERROR:             "Ошибка базы данных: не могу отметить транзакцию как не обрабатывающуюся.",
	INSUFFICIENTBALANCE: "Недостаточный баланс для {{.Purpose}}. Необходимо {{.Sats}}.0f сат больше.",
	TOOSMALLPAYMENT:     "Это слишком мало, пожалуйста, начните {{.Purpose}} от 40 сат.",

	PAYMENTRECEIVED:      "Платёж получен: {{.Sats}}. /tx{{.Hash}}.",
	FAILEDTOSAVERECEIVED: "Платёж получен, но не сохранён в базе данных. Пожалуйста, сообщите о проблеме: <code>{{.Label}}</code>, hash: <code>{{.Hash}}</code>",

	SPAMMYMSG:    "{{if .Spammy}}Теперь эта группа будет спамиться. {{else}}Больше спамить не буду.{{end}}",
	TICKETMSG:    "Новые участники должны заплатить {{.Sat}} сат (убедитесь, что вы установили @{{.BotName}} администратором, чтобы это работало).",
	FREEJOIN:     "К этой группе теперь можно присоединиться свободно.",
	ASKTOCONFIRM: "Заплатить счёт выше?",

	HELPINTRO: `
<pre>{{.Help}}</pre>
Для более подробной информации по каждой команде пожалуйста введите <code>/help &lt;command&gt;</code>.
    `,
	HELPSIMILAR: "/{{.Method}} команда не найдена. Вы имели в виду /{{index .Similar 0}}?{{if gt (len .Similar) 1}} Или может быть /{{index .Similar 1}}?{{if gt (len .Similar) 2}} Возможно {{index .Similar 2}}?{{end}}{{end}}",
	HELPMETHOD: `
<pre>/{{.MainName}} {{.Argstr}}</pre>
{{.Help}}
{{if .HasInline}}
<b>Инлайн запрос</b>
Может быть также вызвана как <a href=\"https://core.telegram.org/bots/inline\">инлайн запрос</a> из группы или в персональном чате, где бот ещё не добавлен. Синтаксис команды похож, но сделан проще: <code>@{{.ServiceId}} {{.InlineExample}}</code>, затем пользователь должен подождать пока результат автоматического "поиска" не будет показан ему.{{end}}
{{if .Aliases}}
<b>Алиасы:</b> <code>{{.Aliases}}</code>{{end}}
    `,

	// the "any" is here only for illustrative purposes. if you call this with 'any' it will
	// actually be assigned to the <сатoshis> variable, and that's how the code handles it.
	RECEIVEHELP: `Создаёт BOLT11 счёт с заданным значением сатоши. Полученные токены будут добавлены к вашему балансу в боте. Если вы не укажете количество, это будет счёт с открытым полем значения, в который может быть добавлено любое количество.",

<code>/receive_320_for_something</code> создаёт запрос на 320 сат с описанием "for something"
<code>/receive 100 за скрытые данные --preimage="0000000000000000000000000000000000000000000000000000000000000000"</code> создаёт счёт с заданным преимаджем (будьте осторожны, вы можете потерять деньги, используйте только если полностью уверены в том, что делаете).
    `,

	PAYHELP: `Расшифровывает BOLT11 счёт и спрашивает хотите ли вы его оплатить (иначе используйте /paynow). Будет получен такой же результат как если бы пользователь скопировал и вставил счёт в чат с ботом. Фото с изображением QR с зашифрованным счётом также работает (если картинка достаточно качественная).

Просто вставляя <code>lnbc1u1pwvmypepp5kjydaerr6rawl9zt7t2zzl9q0rf6rkpx7splhjlfnjr869we3gfqdq6gpkxuarcvfhhggr90psk6urvv5cqp2rzjqtqkejjy2c44jrwj08y5ygqtmn8af7vscwnflttzpsgw7tuz9r407zyusgqq44sqqqqqqqqqqqqqqqgqpcxuncdelh5mtthgwmkrum2u5m6n3fcjkw6vdnffzh85hpr4tem3k3u0mq3k5l3hpy32ls2pkqakpkuv5z7yms2jhdestzn8k3hlr437cpajsnqm</code> расшифровывает и печатает заданный счёт.  
<code>/paynow lnbc1u1pwvmypepp5kjydaerr6rawl9zt7t2zzl9q0rf6rkpx7splhjlfnjr869we3gfqdq6gpkxuarcvfhhggr90psk6urvv5cqp2rzjqtqkejjy2c44jrwj08y5ygqtmn8af7vscwnflttzpsgw7tuz9r407zyusgqq44sqqqqqqqqqqqqqqqgqpcxuncdelh5mtthgwmkrum2u5m6n3fcjkw6vdnffzh85hpr4tem3k3u0mq3k5l3hpy32ls2pkqakpkuv5z7yms2jhdestzn8k3hlr437cpajsnqm</code> оплачивает счёт без подтверждения.
/withdraw_lnurl_3000 создаёт lnurl и QR код для вывода 3000 сатоши из <a href="https://lightning-wallet.com">совместимого кошелька</a> без запроса подтверждения.
/withdraw_lnurl создаёт lnurl и QR код для вывода любого количества, но запросит подтверждение.
<code>/pay</code>, когда отправлено как ответ на другое сообщение, содержащее счёт (например в групповом чате), спрашивает в чате с ботом, согласен ли пользователь оплатить счёт.
    `,

	SENDHELP: `Отправляет сатоши другим пользователям Telegram. Получатель оповещается в его чате с ботом. Если получатель ещё не запустил бот, или заблокировал его, он не может быть оповещён. В этом случае вы можете отменить транзакцию после, вызвав /transactions.

<code>/tip 100</code>, Если эта команда вызывается в ответ на сообщение в групповом чате, где бот добавлен, то автору сообщения будет начислено 100 сатоши.
<code>/send 500 @username</code> отправляет 500 сатоши пользователю Telegram @username.
<code>/send anonymously 1000 @someone</code> то же, что выше, Telegram пользователь @someone увидит только: "Кто-то отправил вам 1000 сатоши".
    `,

	BALANCEHELP: "Покажет вам текущий баланс в сатоши, плюс сумму всего, что вы получили и отправили внутри бота и сумму всех комиссий.",

	GIVEAWAYHELP: `Создаст кнопку в групповом чате. Первый, кто нажмёт на неё, получит сатоши.

/giveaway_1000: как только кто-либо нажмёт 'Получить' 1000 сатоши будут переведены кликеру.
    `,
	GIVEAWAYSATSGIVENPUBLIC: "{{.Sats}} сат подарены от {{.From}} пользователю {{.To}}.{{if .ClaimerHasNoChat}} Для управления своими токенами, начните диалог с @{{.BotName}}.{{end}}",
	CLAIMFAILED:             "Ошибка запроса {{.BotOp}}: {{.Err}}",
	GIVEAWAYCLAIM:           "Получить",
	GIVEAWAYMSG:             "{{.User}} дарит {{.Sats}} сат!",

	COINFLIPHELP: `Запускает честную лотерею подбрасывания монетки с указанным количеством участников. Все платят такое же количество, как было указано в стоимости входа. Победитель получает всё. Токены перемещаются только в тот момент, когда запускается лотерея.

/coinflip_100_5: 5 участников необходимы для старта, победитель получит 500 сатоши (включая его собственные 100, поэтому чистым выигрышем он получит 400 сатоши).
    `,
	COINFLIPWINNERMSG:      "Вы победитель в подбросе монетки с призом {{.TotalSats}} сат. Проигравшие: {{.Senders}}.",
	COINFLIPGIVERMSG:       "Вы проиграли {{.IndividualSats}} в подбросе монетки. Победителем стал {{.Receiver}}.",
	COINFLIPAD:             "Заплатите {{.Sats}} получите шанс выиграть {{.Prize}}! Осталось {{.SpotsLeft}} из {{.MaxPlayers}} спотов!",
	COINFLIPJOIN:           "Присоединиться к розыгрышу!",
	CALLBACKCOINFLIPWINNER: "Победитель: {{.Winner}}",
	COINFLIPOVERQUOTA:      "Вы превысили квоту игры в монетку.",
	COINFLIPRATELIMIT:      "Пожалуйста, подождите 30 минут перед запуском нового раунда монетки.",

	GIVEFLIPHELP: `Начинает раздачу случайным методом, но, вместо подарка первому пользователю, который нажмёт на кнопку, количество будет разыграно между первыми x кликерами.

/giveflip_100_5: 5 участников необходимо 500 сатоши получит победитель от инициатора команды.
    `,
	GIVEFLIPMSG:       "{{.User}} раздаёт {{.Sats}} сат счастливчику из {{.Participants}} участников!",
	GIVEFLIPAD:        "{{.Sats}} будут разданы. Присоединись и получи шанс выиграть! Осталось {{.SpotsLeft}} из {{.MaxPlayers}} мест!",
	GIVEFLIPJOIN:      "Попробую выиграть!",
	GIVEFLIPWINNERMSG: "{{.Sender}} отправил(а) {{.Sats}} сат {{.Receiver}}. Ничего не получили пользователи: {{.Losers}}.{{if .ReceiverHasНетChat}} Для управления своими деньгами начните диалог с @{{.BotName}}.{{end}}",

	FUNDRAISEHELP: `Начинает краудфандинг с заранее определённым количеством участников и вкладом каждого. Если количество участников будет достигнуто, фандрайзинг будет актуализирован. Иначе он будет отменён через несколько часов.

<code>/fundraise 10000 8 @user</code>: Пользователь @user получит 80000 сатоши, если 8 человек присоединятся к компании.
    `,
	FUNDRAISEAD: `
Фандрайзинг {{.Fund}} в пользу {{.ToUser}}!
Необходимо участников: {{.Participants}}
Вклад каждого: {{.Sats}} сат
Присоединились: {{.Registered}}
    `,
	FUNDRAISEJOIN:        "Присоединяюсь!",
	FUNDRAISECOMPLETE:    "Фандрайзинг для {{.Receiver}} завершён!",
	FUNDRAISERECEIVERMSG: "Вы получили {{.TotalSats}} сат от пользователей {{.Senders}}",
	FUNDRAISEGIVERMSG:    "Вы пожертвовали {{.IndividualSats}} в пользу {{.Receiver}}.",

	BLUEWALLETHELP: `Показывает ваши регистрационные данные для импорта кошелька бота в BlueWallet. Вы можете использовать тот же аккаунт из обоих мест попеременно.

/bluewallet Печатает строчку вроде "lndhub://&lt;login&gt;:&lt;password&gt;@&lt;url&gt;", которая должна быть скопирована и вставлена в BlueWallet функцию импорта.
/bluewallet_refresh очищает предыдущий пароль и печатает новую строку. Вы должны ре-импортировать регистрационные данные в кошелёк BlueWallet после этого шага. Делайте это только в том случае, если предыдущие данные были скомпрометированы.
    `,
	BLUEWALLETPASSWORDUPDATEERROR: "Ошибка обновления пароля. Сообщите о ней: {{.Err}}",
	BLUEWALLETCREDENTIALS:         "<code>{{.Credentials}}</code>",

	HIDEHELP: `Скрывает сообщение, которое может быть открыто позже после оплаты. Специальный символ "~" используется для того, чтобы разделить сообщение для предпросмотра ("нажмите здесь, чтобы открыть секрет! ~ это секрет.")

<code>/hide 500 'совершенно секретное сообщение'</code> скрывает "совершенно секретное сообщение" и возвращает id. Позже можно выпустить приглашение к открытию сообщения с помощью /reveal &lt;id_скрытого_сообщения&gt;
<code>/hide 2500 'только богатеи смогут посмотреть это сообщение' ~ 'поздравляю! вы действительно богаты'</code>: в этом случае все потенциальные адресаты скрытого сообщения будут видеть часть перед символом "~" в общем доступе.

Любой пользователь может открыть любое скрытое сообщение (после уплаты), набрав <code>/reveal &lt;id_скрытого_сообщения&gt;</code> в своём приватном чате с ботом, но id известен только создателю сообщения, если он его не разгласил.

Основной способ распространить сообщение -- это кликнуть кнопку "Шарить" и использовать inline запрос в группе или чате. Это действите опубликует предпросмотр в чате так же как и кнопку, которую люди могут нажать и открыть сообщение в приватном чате. Вы можете управлять тем, будет ли сообщение открыто в месте публикации в группе, или приватно только тому, кто оплатит его с помощью флага <code>--public</code>. По умолчанию сообщение приватное.

Вы также можете управлять количеством людей, которые могут просмотреть сообщение приватно с помощью флага <code>--revealers</code> или тем, сколько человек должны будут оплатить показ с помощью флага <code>--crowdfund</code>.

<code>/hide 100 'три человека должны заплатить, чтобы увидеть это сообщение' --crowdfund 3</code>: сообщение будет опубликовано, если 3 человека заплатят по 100 сатоши.
<code>/hide 321 'тольок три человека увидят это сообщение' ~ "вы среди 3 привилегированных человек" --revealers 3</code>: сообщение будет показано приватно только тем 3 людям, которые первыми кликнут на него.
    `,
	REVEALHELP: `Открывает сообщение, которое было скрыто. Автор скрытого сообщения не будет раскрыт. Как только сообщение скрыто, оно может быть открыто глобально, но только теми, кто знает скрытый id.

Приглашение к открытию может быть создано в чате или группе нажатием кнопки "Шэрить" после того, как было скрыто сообщение. Затем применяются обычные правила открытия сообщения, смотрите /help_hide для подробной справки.

<code>/reveal 5c0b2rh4x</code> создаёт приглашение открыть скрытое сообщение 5c0b2rh4x, если оно существует.
    `,
	HIDDENREVEALBUTTON:   `{{.Sats}} открыть {{if .Public}} тут же{{else }} приватно{{end}}. {{if gt .Crowdfund 1}}Краудфандинг: {{.HavePaid}}/{{.Crowdfund}}{{else if gt .Times 0}}Допущены публикаторы: {{.HavePaid}}/{{.Times}}{{end}}`,
	HIDDENDEFAULTPREVIEW: "Здесь скрыто сообщение. {{.Sats}} сат нужно, чтобы его открыть.",
	HIDDENWITHID:         "Сообщение скрыто с id <code>{{.HiddenId}}</code>. {{if gt .Message.Crowdfund 1}}Будет раскрыто публично один раз {{.Message.Crowdfund}} люди заплатят {{.Message.Satoshis}}{{else if gt .Message.Times 0}}Будет раскрыто приватно {{.Message.Times}} пользователям {{else if .Message.Public}}Будет раскрыто публично как только один человек заплатит {{.Message.Satoshis}}{{else}}Будет раскрыто приватно любому заплатившему {{end}}.",
	HIDDENSOURCEMSG:      "Скрытое сообщение <code>{{.Id}}</code> было открыто {{.Revealers}}. Вы получили {{.Sats}} сат.",
	HIDDENREVEALMSG:      "{{.Sats}} сат заплачено для открытия сообщения <code>{{.Id}}</code>.",
	HIDDENSTOREFAIL:      "Не получилось сохранить скрытый контент. Пожалуйста, сообщите об ошибке: {{.Err}}",
	HIDDENMSGNOTFOUND:    "Скрытое сообщение не найдено.",
	HIDDENSHAREBTN:       "Распространить в другом чате",

	BITFLASHCONFIRM:      `<b>[bitflash]</b> Вы подтверждаете запрос транзакции Bitflash на <b>{{.BTCAmount}} BTC</b> на адрес <code>{{.Address}}</code>? Вы заплатите <b>{{printf "%.0f" .Sats}}</b>.`,
	BITFLASHTXQUEUED:     "Транзакция поставлена в очередь!",
	BITFLASHFAILEDTOSAVE: "Ошибка сохранения ордера Bitflash. Пожалуйста, сообщите об ошибке: {{.Err}}",
	BITFLASHLIST: `
<b>[bitflash]</b> Ваши прошлые ордеры
{{range .Orders}}🧱 <code>{{.Amount}}</code> на <code>{{.Address}}</code> <i>{{.Status}}</i>
{{else}}
<i>~ не было ни одного ордера. ~</i>
{{end}}
    `,
	BITFLASHHELP: `
<a href="https://bitflash.club/">Bitflash</a> это сервис, которые осуществляет дешёвые ончейн транзакции при получении платежей Lightning. Он может их делать дёшево, поскольку объединяет несколько транзакций Lightning и затем делает ончейн транзакцию после достижения заданного объёма.

/bitflash_100000_3NRnMC5gVug7Mb4R3QHtKUcp27MAKAPbbJ осуществляет ончейн транзакцию на заданный адрес через bitflash.club общую комиссию. Будет запрошено подтверждение.
/bitflash_orders</code> показывает ваши предыдущие транзакции.
    `,

	MICROBETBETHEADER:           "<b>[Microbet]</b> Сделайте ставки на одно из этих предсказаний:",
	MICROBETINVALIDRESPONSE:     "microbet.fun вернул неправильный ответ.",
	MICROBETPAIDBUTNOTCONFIRMED: "Оплачено, но не подтверждено. Крупная ошибка Microbet?",
	MICROBETPLACING:             "Помещаем ставку на <b>{{.Bet.Description}} ({{if .Back}}🔳{{else}}🔳{{end}})</b>.",
	MICROBETPLACED:              "Ставка размещена!",
	MICROBETFAILEDTOPAY:         "Ошибка оплаты ставки.",
	MICROBETLIST: `
<b>[Microbet]</b> Ваши ставки
{{range .Bets}}<code>{{.Description}}</code> {{if .UserBack}}{{.UserBack}}/{{.Backers}} × {{.Layers}}{{else}}{{.Backers}} × {{.UserLay}}/{{.Layers}}{{end}} <code>{{.Amount}}</code> <i>{{if .Canceled}}отменено{{else if .Closed}}{{if .WonAmount}}выигрыш {{.AmountWon}}{{else}}проигрыш {{.AmountLost}}{{end}}{{else}}открыты{{end}}</i>
{{else}}
<i>~ не было ни одной ставки. ~</i>
{{end}}
    `,
	MICROBETBALANCEERROR: "Ошибка запроса баланса Microbet: {{.Err}}",
	MICROBETBALANCE:      "<b>[Microbet]</b> баланс: <i>{{.Balance}} сат</i>",
	MICROBETHELP: `
<a href="https://microbet.fun/">Microbet</a> это простой сервис, который позволяет ставить на результат спортивных игр. Величина ставки фиксирована и шансы вычисляются с учётом количества ставок. Имеется 1% комиссия на все выводы.

/microbet_bet показывает все открытые ставки, в которых может быть и ваша.
/microbet_bets показывает ваши ставки.
/microbet_balance показывает ваш баланс.
/microbet_withdraw выводит весь ваш баланс.
    `,

	SATELLITEFAILEDTOSTORE:     "Ошибка сохранения заказа на передачу данных. Пожалуйста, сообщите: {{.Err}}",
	SATELLITEFAILEDTOGET:       "Ошибка запроса сохранённых спутниковых данных: {{.Err}}",
	SATELLITEPAID:              "Передача <code>{{.UUID}}</code> оплачена!",
	SATELLITEFAILEDTOPAY:       "Ошибка оплаты передачи.",
	SATELLITEBUMPERROR:         "Ошибка поднятия приоритета передачи: {{.Err}}",
	SATELLITEFAILEDTODELETE:    "Ошибка удаления заказа. Пожалуйста, сообщите: {{.Err}}",
	SATELLITEDELETEERROR:       "Ошибка удаления передачи: {{.Err}}",
	SATELLITEDELETED:           "Передача удалена.",
	SATELLITETRANSMISSIONERROR: "Ошибка создания передачи: {{.Err}}",
	SATELLITEQUEUEERROR:        "Ошибка запроса очереди: {{.Err}}",
	SATELLITEQUEUE: `
<b>[Satellite]</b> Запрошенные передачи
{{range .Orders}}{{.}}
{{else}}
<i>Очередь пуста, всё было передано.</i>
{{end}}
    `,
	SATELLITELIST: `
<b>[Satellite]</b> Ваши спутниковые передачи
{{range .Orders}}{{.}}
{{else}}
<i>Не было спутниковых сообщений.</i>
{{end}}
    `,
	SATELLITEHELP: `
<a href="https://blockstream.com/сатellite/">Blockstream Satellite</a> это сервис, который вещает блоки Bitcoin и другие данные по всему миру. Вы можете передать любые сообщения, заплатив за них сатоши.

<code>/app сатellite 13 'привет со спутника! голосуйте за трампа!'</code> запрашивает передачу через спутник со ставкой 13 сатоши.
/сатellite_transmissions показывает ваши передачи.
    `,

	GOLIGHTNINGFAIL:   "<b>[GoLightning]</b> Ошибка создания ордера: {{.Err}}",
	GOLIGHTNINGFINISH: "<b>[GoLightning]</b> Завершите свой ордер отправкой <code>{{.Order.Price}} BTC</code> на <code>{{.Order.Address}}</code>.",
	GOLIGHTNINGHELP: `
<a href="https://golightning.club/">GoLightning.club</a> это самый дешёвый путь к получению ончейн средств из Lightning сети, всего за 99 сатоши за ордер. В первую очередь вы указываете, сколько вы хотите получить, затем вы отправляете деньги плюс комиссия провайдеру адреса BTC. Готово.

/golightning_1000000 создаёт ордер для перевода 0.01000000 BTC с ончейн адреса на баланс бота.
    `,

	GIFTSHELP: `
<a href="https://lightning.gifts/">Lightning Gifts</a> это лучший способ отослать сатоши в качестве подарка. Простой сервис, простая ссылка URL, нет блокировки средств <b>нет комиссии</b>.

/gifts_1000 создаёт ваучер на 1000 сатоши.
    `,
	GIFTSERROR:      "<b>[gifts]</b> Ошибка: {{.Err}}",
	GIFTSCREATED:    "<b>[gifts]</b> Подарок создан. Для получения просто пройдите по ссылке <code>https://lightning.gifts/redeem/{{.OrderId}}</code>.",
	GIFTSFAILEDSAVE: "<b>[gifts]</b> Ошибка сохранения вашего подарка. Пожалуйста, сообщите: {{.Err}}",
	GIFTSLIST: `
<b>gifts</b>
{{range .Gifts}}- <a href="https://lightning.gifts/redeem/{{.OrderId}}">{{.Amount}}сат</a> {{if .Spent}}затребовано <i>{{.WithdrawDate}}</i> пользователем {{.RedeemerURL}}{{else}}ещё не затребовано{{end}}
{{else}}
<i>~ никаких подарков ещё не было сделано. ~</i>
{{end}}
    `,

	POKERDEPOSITFAIL:  "<b>[Poker]</b> Ошибка депозита: {{.Err}}",
	POKERWITHDRAWFAIL: "<b>[Poker]</b> Ошибка вывода: {{.Err}}",
	POKERBALANCEERROR: "<b>[Poker]</b> Ошибка запроса баланса: {{.Err}}",
	POKERSECRETURL:    `<a href="{{.URL}}">Ваша персональная секретная ссылка на покер здесь, никогда не разглашайте её.</a>`,
	POKERBALANCE:      "<b>[Poker]</b> Баланс: {{.Balance}}",
	POKERSTATUS: `
<b>[Poker]</b>
Игроков онлайн: {{.Players}}
Активных столов: {{.Tables}}
Сатоши в игре: {{.Chips}}

/poker_play чтобы играть здесь!
/poker_url чтобы играть в окне браузера!
    `,
	POKERNOTIFY: `
<b>[Poker]</b> Играют {{.Playing}} человек {{if ne .Waiting 0}}и {{.Waiting}} ожидают игры {{end}}в покер прямо сейчас{{if ne .Sats 0}} , всего {{.Sats}} сатоши в игре{{end}}!

/poker_status для двойной проверки!
/poker_play играть здесь!
/poker_url играть в окне браузера!
    `,
	POKERSUBSCRIBED: "Вы можете играть в покер в течение следующих {{.Minutes}} минут.",
	POKERHELP: `<a href="https://lightning-poker.com/">Lightning Poker</a> первый и простейший безлимитный техасский Холдем Покер, разыгрываемый прямо с использованием сатоши. Просто присоединитесь к столу и начните копить сатоши.

Играя с аккаунта, привязанного к балансу бота, вы можете просто сесть за стол и ваш покерный аккаунт будет пополнен с баланса бота, с минимальными заботами.

/poker_deposit_10000 кладёт 10000 сатоши в ваш стейк.
/poker_balance показывает сколько у вас есть на аккаунте покера.
/poker_withdraw возвращает все деньги обратно на баланс бота.
/poker_status показывает активность покерных столов прямо сейчас.
/poker_url показывает ваш <b>секретный</b> URL, который вы можете открыть из любого браузера и даёт доступ к вашему балансу.
/poker_play показывает виджет игры.
/poker_watch_120 поместит вас в состояние подписки на игры в течение 2 часов и оповестит остальных подписанных людей о том, что вы хотите играть. Вы будете оповещены, если кто-то захочет играть.
    `,

	TOGGLEHELP: `Включает/выключает функции бота в группах. В супергруппах команда может быть запущена только администраторами.

<code>/toggle ticket 10</code> начинает брать комиссию для всех новых пользователей. Полезная функция антиспама. Деньги идут владельцу группы.
<code>/toggle ticket</code> перестаёт брать комиссию с новых участников. 
<code>/toggle spammy</code>: 'спамная' функция выключена по умолчанию. Когда она включена, нотификации о tip командах будут отсылаться в чат, вместо того, чтобы обрабатываться приватно.
    `,

	HELPHELP: "Показывает полную справку или справку о конкретной команде.",

	STOPHELP: "Бот перестаёт отсылать оповещения.",

	CONFIRMINVOICE: `
{{.Sats}} сат ({{dollar .Sats}})
<i>{{.Desc}}</i>
<b>Hash</b>: {{.Hash}}
<b>Node</b>: {{.Node}} ({{.Alias}})
    `,
	FAILEDDECODE: "Ошибка декодирования счёта: {{.Err}}",
	NOINVOICE:    "Счёт не был получен.",
	BALANCEMSG: `
<b>Баланс</b>: {{printf "%.3f" .Sats}} сат ({{dollar .Sats}})
<b>Всего получено</b>: {{printf "%.3f" .Received}} сат
<b>Всего отправлено</b>: {{printf "%.3f" .Sent}} сат
<b>Всего комиссий оплачено</b>: {{printf "%.3f" .Fees}} сат
    `,
	// {{if ne .CoinflipBalance 0}}<b>Баланс Монетки</b>: {{.CoinflipBalance}} сат ({{.CoinflipWins}} выиграно, {{.CoinflipLoses}} проигрыш)
	// {{end}}
	//     `,
	FAILEDUSER: "Не могу распознать получателя.",
	LOTTERYMSG: `
Раунд лотереи стартовал!
Билет на вход: {{.EntrySats}} сат
Всего участников: {{.Participants}}
Приз: {{.Prize}}
Зарегистрировано: {{.Registered}}
    `,
	INVALIDPARTNUMBER:  "Неверное количество участников: {{.Number}}",
	INVALIDAMOUNT:      "Неверное количество: {{.Amount}}",
	USERSENTTOUSER:     "{{.Sats}} сат отправлено {{.User}}{{if .ReceiverHasНетChat}} (не могу уведомить {{.User}} так как он не начал диалога с ботом{{end}}",
	USERSENTYOUSATS:    "{{.User}} отправил вам {{.Sats}} сат{{if .BotOp}} в ходе {{.BotOp}}{{end}}.",
	RECEIVEDSATSANON:   "Кто-то отослал вам {{.Sats}} сат.",
	FAILEDSEND:         "Ошибка отправки: ",
	QRCODEFAIL:         "QR код не был прочитан: {{.Err}}",
	SAVERECEIVERFAIL:   "Ошибка сохранения получателя. Это вероятно баг.",
	CANTSENDNORECEIVER: "Не могу отправить {{.Sats}}. Не хватает получателя!",
	GIVERCANTJOIN:      "Даритель не может присоединиться!",
	CANTJOINTWICE:      "Нельзя участвовать снова!",
	CANTCANCEL:         "У вас нет возможности отменить это.",
	FAILEDINVOICE:      "Ошибка генерации счёта: {{.Err}}",
	ZEROAMOUNTINVOICE:  "Счета с неопределёнными количествами не поддерживаются поскольку они не являются безопасными.",
	INVALIDAMT:         "Неверное количество: {{.Amount}}",
	STOPNOTIFY:         "Оповещения выключены.",
	WELCOME:            "Ваш аккаунт создан.",
	WRONGCOMMAND:       "Не могу понять команду. /help",
	RETRACTQUESTION:    "Вернуть не затребованное поощрение?",
	RECHECKPENDING:     "Перепроверить платёж в обработке?",
	TXNOTFOUND:         "Не могу найти транзакцию {{.HashFirstChars}}.",
	TXINFO: `<code>{{.Txn.Status}}</code> {{.Txn.PeerActionDescription}} в {{.Txn.TimeFormat}} {{if .Txn.IsUnclaimed}}(💤 не затребовано){{end}}
<i>{{.Txn.Description}}</i>{{if not .Txn.TelegramPeer.Valid}}
{{if .Txn.Payee.Valid}}<b>Оплатил</b>: {{.Txn.PayeeLink}} ({{.Txn.PayeeAlias}}){{end}}
<b>Hash</b>: {{.Txn.Hash}}{{end}}{{if .Txn.Preimage.String}}
<b>Preimage</b>: {{.Txn.Preimage.String}}{{end}}
<b>Количество</b>: {{.Txn.Amount}} сат
{{if not (eq .Txn.Status "RECEIVED")}}<b>Комиссия уплачена</b>: {{.Txn.FeeSatoshis}}{{end}}
{{.LogInfo}}
    `,
	TXLIST: `<b>{{if .Offset}}Транзакция от {{.From}} к {{.To}}{{else}}Последние {{.Limit}} транзакций{{end}}</b>
{{range .Transactions}}<code>{{.StatusSmall}}</code> <code>{{.PaddedSatoshis}}</code> {{.Icon}} {{.PeerActionDescription}}{{if not .TelegramPeer.Valid}}<i>{{.Description}}</i>{{end}} <i>{{.TimeFormatSmall}}</i> /tx{{.HashReduced}}
{{else}}
<i>Ещё нет ни одной транзакции</i>
{{end}}
    `,
}
