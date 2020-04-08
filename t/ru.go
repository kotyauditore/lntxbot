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
	TXPENDING:  "Платёж ещё осуществляется, пожалуйста, проверьте позже.",
	TXCANCELED: "Транзакция отменена.",
	UNEXPECTED: "Неожиданная ошибка: пожалуйста, обратитесь в поддержку.",

	CALLBACKWINNER:  "Победитель: {{.Winner}}",
	CALLBACKERROR:   "{{.BotOp}} ошибка{{if .Err}}: {{.Err}}{{else}}.{{end}}",
	CALLBACKEXPIRED: "{{.BotOp}} время истекло.",
	CALLBACKATTEMPT: "Ищу маршрут. /tx_{{.Hash}}",
	CALLBACKSENDING: "Отправляю платёж.",

	INLINEINVOICERESULT:  "Счёт на {{.Sats}} сат.",
	INLINEGIVEAWAYRESULT: "Раздать {{.Sats}} сат",
	INLINEGIVEFLIPRESULT: "Раздаёт {{.Sats}} сат одному из {{.MaxPlayers}} участников",
	INLINECOINFLIPRESULT: "Лотерея с входным платежом {{.Sats}} сат для {{.MaxPlayers}} участников",
	INLINEHIDDENRESULT:   "{{.HiddenId}} ({{if gt .Message.Crowdfund 1}}собрать:{{.Message.Crowdfund}}{{else if gt .Message.Times 0}}прив:{{.Message.Times}}{{else if .Message.Public}}пуб{{else}}прив{{end}}): {{.Message.Content}}",

	LNURLUNSUPPORTED: "Такой тип lnurl не поддерживается.",
	LNURLAUTHSUCCESS: `
lnurl-auth успех!

<b>домен</b>: <i>{{.Host}}</i>
<b>ключ</b>: <i>{{.PublicKey}}</i>
`,
	LNURLPAYPROMPT: `<code>{{.Domain}}</code> ожидает {{if .FixedAmount}}<i>{{.FixedAmount | printf "%.15g"}} сат</i>{{else}}количество между <i>{{.Min | printf "%.15g"}}</i> и <i>{{.Max | printf "%.15g"}} сат</i>{{end}} для оплаты:

{{if .Text}}<code>{{.Text}}</code>{{else if .HTML}}{{.HTML}}{{end}}

{{if not .FixedAmount}}<b>Ответьте с неким количеством для подтверждения.</b>{{end}}
    `,
	LNURLPAYSUCCESS: `<code>{{.Domain}}</code> ответил:

{{.SuccessAction.Description | html}}
{{if eq .SuccessAction.Tag "url"}}<a href="{{.SuccessAction.Data}}">{{.SuccessAction.Data}}</a>{{end}}
    `,
	LNURLPAYMETADATA: `lnurl-pay метаданные:
<b>домен</b>: <i>{{.Domain}}</i>
<b>lnurl</b>: <i>{{.LNURL}}</i>
<b>транзакция</b>: <i>{{.Hash}}</i> /tx_{{.HashFirstChars}}
    `,

	USERALLOWED:       "Счёт оплачен. {{.User}} допущен.",
	SPAMFILTERMESSAGE: "Привет, {{.User}}. У вас 15 минут, чтобы оплатить счёт на {{.Sats}} сат если вы хотите остаться в этой группе:",

	PAYMENTFAILED: "Платёж не состоялся. /log_{{.ShortHash}}",
	PAIDMESSAGE: `Оплачено <b>{{printf "%.15g" .Sats}} сат</b> (+ {{.Fee}} комиссия). 

<b>Hash:</b> {{.Hash}}{{if .Preimage}}
<b>Proof:</b> {{.Preimage}}{{end}}

/tx_{{.ShortHash}}`,
	OVERQUOTA:           "Вы превысили квоту {{.App}}.",
	RATELIMIT:           "Пожалуйста, подождите 30 минут.",
	DBERROR:             "Ошибка базы данных: не могу отметить транзакцию как не обрабатывающуюся.",
	INSUFFICIENTBALANCE: "Недостаточный баланс для {{.Purpose}}. Необходимо на {{.Sats}}.0f сат больше.",

	PAYMENTRECEIVED:      "Платёж получен: {{.Sats}}. /tx_{{.Hash}}.",
	FAILEDTOSAVERECEIVED: "Платёж получен, но не сохранён в базе данных. Пожалуйста, сообщите о проблеме: <code>{{.Label}}</code>, hash: <code>{{.Hash}}</code>",

	SPAMMYMSG:           "{{if .Spammy}}Теперь эта группа будет спамиться. {{else}}Больше спамить не буду.{{end}}",
	COINFLIPSENABLEDMSG: "Подбросы монетки {{if .Enabled}}разрешены{{else}}запрещены{{end}} в этой группе.",
	LANGUAGEMSG:         "Установлен язык этого чата <code>{{.Language}}</code>.",
	TICKETMSG:           "Новые участники должны заплатить {{.Sat}} сат (убедитесь, что вы установили @{{.BotName}} администратором, чтобы это работало).",
	FREEJOIN:            "К этой группе теперь можно присоединиться свободно.",

	APPBALANCE: "<b>[{{.App}}]</b> Баланс: <i>{{.Balance}} sat</i>",

	HELPINTRO: `
<pre>{{.Help}}</pre>
Для более подробной информации по каждой команде пожалуйста введите <code>/help &lt;command&gt;</code>.
    `,
	HELPSIMILAR: "/{{.Method}} команда не найдена. Вы имели в виду /{{index .Similar 0}}?{{if gt (len .Similar) 1}} Или может быть /{{index .Similar 1}}?{{if gt (len .Similar) 2}} Вероятно, {{index .Similar 2}}?{{end}}{{end}}",
	HELPMETHOD: `
<pre>/{{.MainName}} {{.Argstr}}</pre>
{{.Help}}
{{if .HasInline}}
<b>Инлайн запрос</b>
Может быть вызвана как <a href="https://core.telegram.org/bots/inline">инлайн запрос</a> из группы или в персональном чате, где бот ещё не добавлен. Синтаксис команды проще: <code>@{{.ServiceId}} {{.InlineExample}}</code>, пользователь должен подождать результат "поиска" команды.{{end}}
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

<code>/pay</code>, когда отправлено как ответ на другое сообщение, содержащее счёт (например в групповом чате), спрашивает в чате с ботом, согласен ли пользователь оплатить счёт.
    `,

	SENDHELP: `Отправляет сатоши другим пользователям Telegram. Получатель оповещается в его чате с ботом. Если получатель ещё не запустил бот, или заблокировал его, он не может быть оповещён. В этом случае вы можете отменить транзакцию после, вызвав /transactions.

<code>/tip 100</code>, Если эта команда вызывается в ответ на сообщение в групповом чате, где бот добавлен, то автору сообщения будет начислено 100 сатоши.
<code>/send 500 @username</code> отправляет 500 сатоши пользователю Telegram @username.
<code>/send anonymously 1000 @someone</code> то же, что выше, Telegram пользователь @someone увидит только: "Кто-то отправил вам 1000 сатоши".
    `,

	TRANSACTIONSHELP: `
Показывает все ваши транзакции постранично. Каждая транзакция имеет ссылку, на которую можно нажать, чтобы просмотреть её более подробно.

/transactions показывает все транзакции, начиная с недавних.
<code>/transactions --in</code> показывает только входящие транзакции.
<code>/transactions --out</code> показывает только исходящие транзакции.
    `,

	BALANCEHELP: "Покажет вам текущий баланс в сатоши, а также сумму всего, что вы получили и отправили внутри бота и сумму всех комиссий.",

	GIVEAWAYHELP: `Создаст кнопку в групповом чате. Первый, кто нажмёт на неё, получит сатоши.

/giveaway_1000: как только кто-либо нажмёт 'Получить' 1000 сатоши будут переведены кликеру.
    `,
	GIVEAWAYSATSGIVENPUBLIC: "{{.Sats}} сат подарены от {{.From}} пользователю {{.To}}.{{if .ClaimerHasNoChat}} Для управления своими сатоши, начните диалог с @{{.BotName}}.{{end}}",
	CLAIMFAILED:             "Ошибка запроса {{.BotOp}}: {{.Err}}",
	GIVEAWAYCLAIM:           "Получить",
	GIVEAWAYMSG:             "{{.User}} раздаёт {{.Sats}} сат!",

	COINFLIPHELP: `Запускает честную лотерею подбрасывания монетки с указанным количеством участников. Все платят одинаковую стоимость входа. Победитель получает всё. Токены перемещаются только в тот момент, когда запускается лотерея.

/coinflip_100_5: 5 участников необходимы, победитель получит 500 сатоши (включая его 100, поэтому чистый выигрыш 400 сатоши).`,
	COINFLIPWINNERMSG:      "Вы победитель в подбросе монетки с призом {{.TotalSats}} сат. Проигравшие: {{.Senders}}.",
	COINFLIPGIVERMSG:       "Вы проиграли {{.IndividualSats}} в подбросе монетки. Победителем стал {{.Receiver}}.",
	COINFLIPAD:             "Заплатите {{.Sats}} сат и получите шанс выиграть {{.Prize}}! Осталось {{.SpotsLeft}} из {{.MaxPlayers}} свободных мест!",
	COINFLIPJOIN:           "Присоединиться к розыгрышу!",
	CALLBACKCOINFLIPWINNER: "Победитель: {{.Winner}}",

	GIVEFLIPHELP: `Начинает раздачу случайным методом, но, вместо подарка первому пользователю, который нажмёт на кнопку, количество будет разыграно между первыми x кликерами.

/giveflip_100_5: 5 участников необходимо 500 сатоши получит победитель от инициатора команды.
    `,
	GIVEFLIPMSG:       "{{.User}} раздаёт {{.Sats}} сат счастливчику из {{.Participants}} участников!",
	GIVEFLIPAD:        "{{.Sats}} будут разданы. Присоединись и получи шанс выиграть! Осталось {{.SpotsLeft}} из {{.MaxPlayers}} мест!",
	GIVEFLIPJOIN:      "Попробую выиграть!",
	GIVEFLIPWINNERMSG: "{{.Sender}} отправил(а) {{.Sats}} сат {{.Receiver}}. Ничего не получили пользователи: {{.Losers}}.{{if .ReceiverHasNoChat}} Для управления своими деньгами начните диалог с @{{.BotName}}.{{end}}",

	FUNDRAISEHELP: `Начинает краудфандинг с заранее определённым количеством участников и вкладом каждого. Если количество участников будет достигнуто, фандрайзинг будет актуализирован. Иначе он будет отменён через несколько часов.

<code>/fundraise 10000 8 @user</code>: Пользователь @user получит 80000 сатоши, если 8 человек присоединятся к компании.
    `,
	FUNDRAISEAD: `
Фандрайзинг {{.Fund}} в пользу {{.ToUser}}!
Необходимо участников: {{.Participants}}
Вклад каждого: {{.Sats}} сат
Присоединились: {{.Registered}}
    `,
	FUNDRAISEJOIN:        "Вкладываюсь!",
	FUNDRAISECOMPLETE:    "Фандрайзинг для {{.Receiver}} завершён!",
	FUNDRAISERECEIVERMSG: "Вы получили {{.TotalSats}} сат от пользователей {{.Senders}}",
	FUNDRAISEGIVERMSG:    "Вы пожертвовали {{.IndividualSats}} в пользу {{.Receiver}}.",

	BLUEWALLETHELP: `Показывает ваши регистрационные данные для импорта кошелька бота в BlueWallet. Вы можете использовать тот же аккаунт из обоих мест попеременно.

/bluewallet Печатает строчку вроде "lndhub://&lt;login&gt;:&lt;password&gt;@&lt;url&gt;", которая должна быть скопирована и вставлена в BlueWallet функцию импорта.
/bluewallet_refresh очищает предыдущий пароль и печатает новую строку. Вы должны ре-импортировать регистрационные данные в кошелёк BlueWallet после этого шага. Делайте это только в том случае, если предыдущие данные были скомпрометированы.
    `,
	APIPASSWORDUPDATEERROR: "Ошибка обновления пароля. Сообщите о ней: {{.Err}}",
	APICREDENTIALS: `
Это токены для <i>Basic Auth</i>. API совместимо с lndhub.io с добавленными методами.

Полный доступ: <code>{{.Full}}</code>
Выписка счетов: <code>{{.Invoice}}</code>
Доступ на чтение: <code>{{.ReadOnly}}</code>
API Base URL: <code>{{.ServiceURL}}/</code>

/api_full, /api_invoice и /api_readonly будут показывать эти специфические токены вместе с QR кодами.
/api_url покажет QR код для API Base URL.

Сохраняйте эти токены в секрете. Если они будут скомпрометированы, вызывайте команду /api_refresh для их полной замены.
    `,
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
	HIDDENREVEALBUTTON:   `{{.Sats}} сат открыть {{if .Public}}тут же{{else}}приватно{{end}}. {{if gt .Crowdfund 1}}Краудфандинг: {{.HavePaid}}/{{.Crowdfund}}{{else if gt .Times 0}}Допущены публикаторы: {{.HavePaid}}/{{.Times}}{{end}}`,
	HIDDENDEFAULTPREVIEW: "Здесь скрыто сообщение. {{.Sats}} сат нужно, чтобы его открыть.",
	HIDDENWITHID:         "Сообщение скрыто с id <code>{{.HiddenId}}</code>. {{if gt .Message.Crowdfund 1}}Будет раскрыто публично один раз {{.Message.Crowdfund}} люди заплатят {{.Message.Satoshis}}{{else if gt .Message.Times 0}}Будет раскрыто приватно {{.Message.Times}} пользователям {{else if .Message.Public}}Будет раскрыто публично как только один человек заплатит {{.Message.Satoshis}}{{else}}Будет раскрыто приватно любому заплатившему {{end}}.",
	HIDDENSOURCEMSG:      "Скрытое сообщение <code>{{.Id}}</code> было открыто {{.Revealers}}. Вы получили {{.Sats}} сат.",
	HIDDENREVEALMSG:      "{{.Sats}} сат заплачено для открытия сообщения <code>{{.Id}}</code>.",
	HIDDENMSGNOTFOUND:    "Скрытое сообщение не найдено.",
	HIDDENSHAREBTN:       "Поделиться в другом чате",

	MICROBETBETHEADER:           "<b>[Microbet]</b> Сделайте ставки на одно из этих предсказаний:",
	MICROBETPAIDBUTNOTCONFIRMED: "<b>[Microbet]</b> Оплачено, но не подтверждено. Крупная ошибка Microbet?",
	MICROBETPLACING:             "<b>[Microbet]</b> Размещаем ставку на <b>{{.Bet.Description}} ({{if .Back}}back{else}}lay{{end}})</b>.",
	MICROBETPLACED:              "<b>[Microbet]</b> Ставка размещена!",
	MICROBETLIST: `
<b>[Microbet]</b> Ваши ставки
{{range .Bets}}<code>{{.Description}}</code> {{if .UserBack}}{{.UserBack}}/{{.Backers}} × {{.Layers}}{{else}}{{.Backers}} × {{.UserLay}}/{{.Layers}}{{end}} <code>{{.Amount}}</code> <i>{{if .Canceled}}отменено{{else if .Closed}}{{if .WonAmount}}выигрыш {{.AmountWon}}{{else}}проигрыш {{.AmountLost}}{{end}}{{else}}открыты{{end}}</i>
{{else}}
<i>~ не было ни одной ставки. ~</i>
{{end}}
    `,
	MICROBETHELP: `
<a href="https://microbet.fun/">Microbet</a> это простой сервис, который позволяет ставить на результат спортивных игр. Величина ставки фиксирована и шансы вычисляются с учётом количества ставок. Имеется 1% комиссия на все выводы.

/microbet_bet показывает все открытые ставки, в которых может быть и ваша.
/microbet_bets показывает ваши ставки.
/microbet_balance показывает ваш баланс.
/microbet_withdraw выводит весь ваш баланс.
    `,
	BITREFILLINVENTORYHEADER: `<b>[Bitrefill]</b> Выберите провайдера услуг:`,
	BITREFILLPACKAGESHEADER:  `<b>[Bitrefill]</b> Выберите вашу <i>{{.Item}}</i> карту{{if .ReplyCustom}} (or reply with a custom value){{end}}:`,
	BITREFILLNOPROVIDERS:     `<b>[Bitrefill]</b> Провайдер не найден.`,
	BITREFILLCONFIRMATION:    `<b>[Bitrefill]</b> Действительно купить <i>{{.Package.Value}} {{.Item.Currency}}</i> карту <b>{{.Item.Name}}</b> за <i>{{.Sats}} сат</i> ({{dollar .Sats}})?`,
	BITREFILLFAILEDSAVE:      "<b>[Bitrefill]</b> Ваш заказ <code>{{.OrderId}}</code> был оплачен, но не сохранён. Сообщите об ошибке: {{.Err}}",
	BITREFILLPURCHASEDONE: `<b>[Bitrefill]</b> Ваш заказ <code>{{.OrderId}}</code> был оформлен успешно.
{{if .Info.LinkInfo}}
Ссылка: <a href="{{.Info.LinkInfo.Link}}">{{.Info.LinkInfo.Link}}</a>
Инструкции: <i>{{.Info.LinkInfo.Other}}</i>
{{else if .Info.PinInfo}}
PIN: <code>{{.Info.PinInfo.Pin}}</code>
Инструкции: <i>{{.Info.PinInfo.Instructions}}</i>
<i>{{.Info.PinInfo.Other}}</i>
{{end}}
    `,
	BITREFILLPURCHASEFAILED: "<b>[Bitrefill]</b> Ваш заказ был оплачен, но Bitrefill вернул ошибку, когда его выполнил: <i>{{.ErrorMessage}}</i>. Пожалуйста, сообщите об этой ошибке, чтобы мы могли обратиться в Bitrefill.",
	BITREFILLCOUNTRYSET:     "<b>[Bitrefill]</b> Страна выбрана, {{if .CountryCode}}<code>{{.CountryCode}}</code>{{else}}none{{end}}.",
	BITREFILLINVALIDCOUNTRY: "<b>[Bitrefill]</b> Неверный код страны <code>{{.CountryCode}}</code>. Вы можете выбрать из кодов {{range .Available}} <code>{{.}}</code>{{end}}.",
	BITREFILLHELP: `
<a href="https://www.bitrefill.com/">Bitrefill</a> это крупнейший магазин подарочных сертификатов и сервис пополнения мобильных телефонов, который работает на Bitcoin Lightning Network. Если вы желаете купить реальные товары или услуги с помощью сатоши Lightning, он может быть вашим первым кандидатом.

Для покупки подарочной карты используйте команду /bitrefill после которой укажите название сервиса, который вам интересен. Для пополнения счёта мобильного телефона, сделайте то же самое, но также добавьте номер телефона с кодом страны. Опционально, вы можете также установить свою страну с помощью <code>/bitrefill country</code>, в таком случае вы будете получать только те предложения, которые актуальны для вашей страны и вам не потребуется пропускать стопку разных сервисов Амазона, например.

<code>/bitrefill country RU</code> установит вашу страну по умолчанию как Россия.
<code>/bitrefill country ''</code> сбросит вашу страну по умолчанию.
<code>/bitrefill mts +7411971732181</code> покажет опции пополнения для заданного номера оператора МТС.
<code>/bitrefill amazon</code> покажет опции сертификатов различных номиналов, которые вы можете приобрести для покупок на Амазоне.

Вы также можете найти доступные опции подарочных карт на сайте <a href="https://www.bitrefill.com/">Bitrefill</a>, пройдя по ссылке. Зачастую вебсайт имеет больше предложений и более гибок в выборе карт. При этом, цены будут одинаковы в боте и на сайте.
    `,
	SATELLITEFAILEDTOSTORE:     "Ошибка сохранения заказа на передачу данных. Пожалуйста, сообщите: {{.Err}}",
	SATELLITEFAILEDTOGET:       "Ошибка запроса сохранённых спутниковых данных: {{.Err}}",
	SATELLITEPAID:              "Передача <code>{{.UUID}}</code> оплачена!",
	SATELLITEFAILEDTOPAY:       "Ошибка оплаты передачи.",
	SATELLITETRANSMISSIONERROR: "Ошибка создания передачи: {{.Err}}",
	SATELLITELIST: `
<b>[Satellite]</b> Ваши спутниковые передачи
{{range .Orders}}📡 <code>{{.UUID}}</code> <i>{{.Status}}</i> <code>{{.MessageSize}}b</code> <code>{{printf "%.62f" .BidPerByte}} msat/b</code> <i>{{.Time}}</i>
{{else}}
<i>Не было спутниковых сообщений.</i>
{{end}}
    `,
	SATELLITEHELP: `
<a href="https://blockstream.com/satellite/">Blockstream Satellite</a> это сервис, который вещает блоки Bitcoin и другие данные по всему миру. Вы можете передать любые сообщения, заплатив за них сатоши.

<code>/app satellite 13 'привет со спутника! голосуйте за трампа!'</code> запрашивает передачу через спутник со ставкой 13 сатоши.
/satellite_transmissions показывает ваши передачи.
    `,

	FUNDBTCFINISH: "Завершите свой ордер отправкой <code>{{.Order.Price}} BTC</code> на <code>{{.Order.Address}}</code>.",
	FUNDBTCHELP: `
<a href="https://golightning.club/">golightning.club</a> это самый дешёвый путь к получению ончейн средств из Lightning сети, всего за 99 сатоши за ордер. В первую очередь вы указываете, сколько вы хотите получить, затем вы отправляете деньги плюс комиссия провайдеру адреса BTC. Готово.

/fundbtc_1000000 создаёт ордер для перевода 0.01000000 BTC с ончейн адреса на баланс бота.
    `,
	BITCLOUDSHELP: `
<a href="https://bitclouds.sh">bitclouds.sh</a> это программируемая платформа VPS, специализированная на Биткоин-сервисах. Вы можете получить обычный VPS, предназначенный для Bitcoin Core или инстансы c-lightning за 66 сат в час. Не существует более дешёвого сервиса!

/bitclouds отразит статус активных хостов.
/bitclouds_create отобразит доступные опции создания хоста.
<code>/bitclouds topup &lt;sats&gt;</code> пополнит счёт хоста или позволит выбрать определённый хост.

Также @{{.BotName}} будет напоминать вам о пополнении счёта хостов, как только на них будет снижаться баланс.
    `,
	BITCLOUDSCREATEHEADER: "<b>[bitclouds]</b> Выберите хост:",
	BITCLOUDSCREATED: `<b>[bitclouds]</b> Ваш <i>{{.Image}}</i> хост <code>{{.Host}}</code> готов!
{{with .Status}}
  {{if .SSHPwd}}<b>ssh доступ:</b>
  <pre>ssh-copy-id -p{{.SSHPort}} {{.SSHUser}}@{{.IP}}
# type password: {{.SSHPwd}}
ssh -p{{.SSHPort}} {{.SSHUser}}@{{.IP}}</pre>{{end}}
  {{with .Sparko}}<b>Посетите ваш <a href="{{.}}">Спарк-кошелёк</a> или вызовите c-lightning RPC снаружи:</b>
<b>Обратиться к c-lightning RPC снаружи:</b>
  <pre>curl -kX POST {{.}}/rpc -d '{"method": "getinfo"}' -H 'X-Access: grabyourkeyinside'</pre>{{end}}
  {{if .RPCPwd}}<b>Вызвать Bitcoin Core RPC:</b>
  <pre>bitcoin-cli -rpcport={{.RPCPort}} -rpcuser={{.RPCUser}} -rpcpassword={{.RPCPwd}} getblockchaininfo</pre>{{end}}
  Hours left in balance: <b>{{.HoursLeft}}</b>
{{end}}
    `,
	BITCLOUDSSTOPPEDWAITING: "<b>[bitclouds]</b> Таймаут во время ожидания подготовки хоста bitclouds.sh, вызовите /bitclouds_status_{{.EscapedHost}} в течение пары минут, если он по-прежнему не будет доступен, сообщите о проблеме, используя доказательство платежа.",
	BITCLOUDSNOHOSTS:        "<b>[bitclouds]</b> На вашем аккаунте нет хостов. Может быть, вы желаете создать /bitclouds_create?",
	BITCLOUDSHOSTSHEADER:    "<b>[bitclouds]</b> Выберите ваш хост:",
	BITCLOUDSSTATUS: `<b>[bitclouds]</b> Хост <code>{{.Host}}</code>:
{{with .Status}}
  Статус: <i>Подписан</i>
  Баланс: <i>{{.HoursLeft}} часов осталось</i>
  IP: <code>{{.IP}}</code>
  {{if .UserPort }}App порт: <code>{{.UserPort}}</code>
  {{end}}{{if .SSHPort}}SSH: <code>ssh -p{{.SSHPort}} {{.SSHUser}}@{{.IP}}</code>
  {{end}}{{with .Sparko}}<a href="{{.}}">Sparko</a>: <code>curl -k -X POST {{.}}/rpc -d '{"method": "getinfo"}' -H 'X-Access: grabyourkeyinside'</code>
  {{end}}{{if .RPCPwd}}Bitcoin Core: <code>bitcoin-cli -rpcconnect={{.IP}} -rpcport={{.RPCPort}} -rpcuser={{.RPCUser}} -rpcpassword={{.RPCPwd}} getblockchaininfo</code>
  {{end}}
{{end}}
    `,
	BITCLOUDSREMINDER: `<b>[bitclouds]</b> {{if .Alarm}}⚠{{else}}⏰{{end}} Bitclouds хост <code>{{.Host}}</code> израсходует средства в ближайшие {{if .Alarm}}<b>{{.TimeToExpire}}</b> и <i>всё будет удалено</i>!{{else}}{{.TimeToExpire}}.{{end}}

{{if .Alarm}}⚠⚠⚠⚠⚠

{{end}}Используйте /bitclouds_topup_{{.Sats}}_{{.EscapedHost}} для оплаты следующей недели!
    `,

	QIWIHELP: `
Мгновенно переводит ваши сатоши на кошелёк <a href="https://qiwi.com/">Qiwi</a>. Сервис осуществляется @lntorubbot.

<code>/qiwi 50 rub to 777777777</code> отправляет эквивалент 50 рублей на номер 77777777.
<code>/qiwi default 999999999</code> устанавливает номер 999999999 как ваш аккаунт по умолчанию.
<code>/qiwi 10000 sat</code> отправляет 10000 сат с конвертацией в рубли на ваш аккаунт.
/qiwi_default показывает ваш аккаунт по умолчанию.
/qiwi_list показывает ваши прошлые транзакции.
    `,
	YANDEXHELP: `
Мгновенно переводит ваши сатоши на аккаунт <a href="https://money.yandex.ru/">Yandex.Money</a>. Сервис осуществляется @lntorubbot.

<code>/yandex 50 rub to 777777777</code> отправляет эквивалент 50 рублей на номер 77777777.
<code>/yandex default 999999999</code> устанавливает номер 999999999 как ваш аккаунт по умолчанию.
<code>/yandex 10000 sat</code> отправляет 10000 сат с конвертацией в рубли на ваш аккаунт.
/yandex_default показывает ваш аккаунт по умолчанию.
/yandex_list показывает ваши прошлые транзакции.
    `,
	LNTORUBCONFIRMATION:  "Отправляю <i>{{.Sat}} сат ({{.Rub}} рублей)</i> на <b>{{.Type}}</b> счёт <code>{{.Target}}</code>. Подтверждаете?",
	LNTORUBFULFILLED:     "<b>[{{.Type}}]</b> Перевод <code>{{.OrderId}}</code> завершён.",
	LNTORUBCANCELED:      "<b>[{{.Type}}]</b> Перевод <code>{{.OrderId}}</code> отменён.",
	LNTORUBFIATERROR:     "<b>[{{.Type}}]</b> Ошибка отправки рублей. Пожалуйста, сообщите об этой проблеме, упомянув идентификатор <code>{{.OrderId}}</code>.",
	LNTORUBMISSINGTARGET: "<b>[{{.Type}}]</b> Вы не указали получателя и получатель по умолчанию не установлен!",
	LNTORUBDEFAULTTARGET: `<b>[{{.Type}}]</b> Обмен по умолчанию: {{.Target}}`,
	LNTORUBORDERLIST: `<b>[{{.Type}}]</b>
{{range .Orders}}<i>{{.Sat}} сат ({{.Rub}} руб)</i> на <code>{{.Target}}</code> время <i>{{.Time}}</i>
{{else}}
<i>~ не было ни одного обмена. ~</i>
{{end}}
    `,

	GIFTSHELP: `
<a href="https://lightning.gifts/">Lightning Gifts</a> это лучший способ отослать сатоши в качестве подарка. Простой сервис, простая ссылка URL, нет блокировки средств <b>нет комиссии</b>.

Создавая свои сатоши-купоны в @{{ .BotName }} вы можете отслеживать их погашение и понять какие не были востребованы.

/gifts показывает купоны, которые вы создали.
/gifts_1000 создаёт ваучер на 1000 сатоши.
    `,
	GIFTSCREATED:    "<b>[gifts]</b> Подарок создан. Для получения просто пройдите по ссылке <code>https://lightning.gifts/redeem/{{.OrderId}}</code>.",
	GIFTSFAILEDSAVE: "<b>[gifts]</b> Ошибка сохранения вашего подарка. Пожалуйста, сообщите: {{.Err}}",
	GIFTSLIST: `
<b>gifts</b>
{{range .Gifts}}- <a href="https://lightning.gifts/redeem/{{.OrderId}}">{{.Amount}}сат</a> {{if .Spent}}затребовано <i>{{.WithdrawDate}}</i> пользователем {{.RedeemerURL}}{{else}}ещё не затребовано{{end}}
{{else}}
<i>~ никаких подарков ещё не было сделано. ~</i>
{{end}}
    `,
	GIFTSSPENTEVENT: `<b>[gifts]</b> Купон использован!

Ваши {{.Amount}} сат <code>{{.Id}}</code> был востребованы {{if .Description}} из счёта с описанием (отметкой)
<i>{{.Description}}</i>{{end}}.
    `,

	TOGGLEHELP: `Включает/выключает функции бота в группах. В супергруппах команда может быть запущена только администраторами.

/toggle_ticket_10  начинает брать комиссию для всех новых пользователей. Полезная функция антиспама. Деньги идут владельцу группы.
/toggle_ticket перестаёт брать комиссию с новых участников.
/toggle_language_ru меняет язык чата на Русский, 
/toggle_language показывает язык чата, это также работает в приватном чате бота.
/toggle_spammy 'спамная' функция выключена по умолчанию. Когда она включена, нотификации о tip командах будут отсылаться в чат, вместо того, чтобы обрабатываться приватно.
`,
	SATS4ADSHELP: `
Sats4ads это маркетплейс рекламы в Telegram. Платите сатоши, чтобы показывать рекламу остальным, получайте сатоши за каждое рекламное объявление, которое вы видите.

Цены для каждого пользователя устанавливаются в миллисатоши за символ.
Каждое объявление также включает фиксированную плату в 1 сат.
Картинки и видео эквивалентны 300 символам.

/sats4ads_on_15 устанавливает аккаунт в режим просмотра рекламы. Любой может опубликовать сообщение для вас за 15 миллисатоши за символ. Вы можете изменить эту цену
/sats4ads_off выключает подписку на рекламные сообщения.
/sats4ads_rates показывает разбиение пользователей по ценовым группам. Полезно, чтобы планировать бюджет.
/sats4ads_broadcast_1000 вещает сообщение. Последняя цифра равна максимуму сатоши, которые будут потрачены. Более дешёвые подписчики рекламы предпочтительны. Эта команда должна вызываться в ответ на другое сообщение, содержание которого будет использовано в качестве текста рекламы.
    `,
	SATS4ADSTOGGLE:    `<b>[sats4ads]</b> {{if .On}}Смотреть рекламу и получать {{printf "%.15g" .Sats}} сат за символ.{{else}}Вы больше не увидите рекламы.{{end}}`,
	SATS4ADSBROADCAST: `<b>[sats4ads]</b> {{if .NSent}}Сообщение отправлено {{.NSent}} раз с полной стоимостью {{.Sats}} сат ({{dollar .Sats}}).{{else}}Не могу найти подписчиков с подходящими параметрами. /sats4ads_rates{{end}}`,
	SATS4ADSPRICETABLE: `<b>[sats4ads]</b> Количество пользователей в каждом диапазоне цены.
{{range .Rates}}<code>{{.UpToRate}} мсат</code>: <i>пользователей всего {{.NUsers}}</i>
{{else}}
<i>Никто ещё не зарегистрировался на просмотры объявлений.</i>
{{end}}
Каждое сообщение стоит такую цену <i>за символ</i> + <code>1 сат</code> за каждого пользователя.
    `,
	SATS4ADSADFOOTER: `[sats4ads: {{printf "%.15g" .Sats}} сат]`,

	HELPHELP: "Показывает полную справку или справку о конкретной команде.",

	STOPHELP: "Бот перестаёт отсылать оповещения.",

	PAYPROMPT: `
{{if .Sats}}{{.Sats}} сат ({{dollar .Sats}})
{{end}}<i>{{.Desc}}</i>
<b>Hash</b>: {{.Hash}}
<b>Node</b>: {{.Node}} ({{.Alias}})

{{if .Sats}}Заплатить счёт выше?
{{else}}<b>Ответьте с желаемым количеством для подтверждения</b>
{{end}}
    `,
	FAILEDDECODE: "Ошибка декодирования счёта: {{.Err}}",
	BALANCEMSG: `
<b>Полный баланс</b>: {{printf "%.15g" .Sats}} сат ({{dollar .Sats}})
<b>Доступный баланс</b>: {{printf "%.15g" .Sats}} сат ({{dollar .Usable}})
<b>Всего получено</b>: {{printf "%.15g" .Received}} сат
<b>Всего отправлено</b>: {{printf "%.15g" .Sent}} сат
<b>Всего комиссий оплачено</b>: {{printf "%.15g" .Fees}} сат

/balance_apps
/transactions
    `,
	TAGGEDBALANCEMSG: `
<b>Всего разница </b> <code>получено - потрачено</code> <b>на внутренние и внешние</b> /apps<b>:</b>

{{range .Balances}}<code>{{.Tag}}</code>: <i>{{printf "%.15g" .Balance}} сат</i>  ({{dollar .Balance}})
{{else}}
<i>Пока не совершено транзакций данного типа.</i>
{{end}}
    `,
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
	USERSENTTOUSER:     "{{.Sats}} сат отправлено {{.User}}{{if .ReceiverHasNoChat}} (не могу уведомить {{.User}} так как он не начал диалога с ботом{{end}}",
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
	INVALIDAMT:         "Неверное количество: {{.Amount}}",
	STOPNOTIFY:         "Оповещения выключены.",
	WELCOME: `
Добро пожаловать. Ваш аккаунт создан. С этого момента вы можете пополнить его биткоинами, изнутри Telegram. Пожалуйста, помните, что мы не можем гарантировать сохранность ваших средств в случае, если мы потеряем средства из-за программной ошибки или атаки хакеров. Не храните на балансе бота больше, чем вы можете позволить себе потерять.

Помимо этого, бот достаточно безопасен.

По любым имеющимся вопросам и, чтобы просто поприветствовать, присоединяйтесь в чат @lntxbot_dev (предупреждение: там может быть установлена входная плата в сатоши).
    `,
	WRONGCOMMAND:    "Не могу понять команду. /help",
	RETRACTQUESTION: "Вернуть не затребованное поощрение?",
	RECHECKPENDING:  "Перепроверить платёж в обработке?",
	TXNOTFOUND:      "Не могу найти транзакцию {{.HashFirstChars}}.",
	TXINFO: `<code>{{.Txn.Status}}</code> {{.Txn.PeerActionDescription}} в {{.Txn.TimeFormat}} {{if .Txn.IsUnclaimed}}(💤 не затребовано){{end}}
<i>{{.Txn.Description}}</i>{{if not .Txn.TelegramPeer.Valid}}
{{if .Txn.Payee.Valid}}<b>Оплатил</b>: {{.Txn.PayeeLink}} ({{.Txn.PayeeAlias}}){{end}}
<b>Hash</b>: {{.Txn.Hash}}{{end}}{{if .Txn.Preimage.String}}
<b>Preimage</b>: {{.Txn.Preimage.String}}{{end}}
<b>Количество</b>: {{.Txn.Amount | printf "%.15g"}} сат
{{if not (eq .Txn.Status "RECEIVED")}}<b>Комиссия уплачена</b>: {{.Txn.FeeSatoshis}}{{end}}
{{.LogInfo}}
    `,
	TXLIST: `<b>{{if .Offset}}Транзакция от {{.From}} к {{.To}}{{else}}Последние {{.Limit}} транзакций{{end}}</b>
{{range .Transactions}}<code>{{.StatusSmall}}</code> <code>{{.PaddedSatoshis}}</code> {{.Icon}} {{.PeerActionDescription}}{{if not .TelegramPeer.Valid}}<i>{{.Description}}</i>{{end}} <i>{{.TimeFormatSmall}}</i> /tx_{{.HashReduced}}
{{else}}
<i>Ещё нет ни одной транзакции</i>
{{end}}
    `,
	TUTORIALWALLET: `
@{{.BotName}} это Lightning кошелёк, который работает из вашего аккаунта Telegram.

Вы можете использовать его для игр и получения счетов Lightning, он сохраняет ваш баланс и историю ваших транзакций.

Он также поддерживает <a href="https://github.com/btcontract/lnurl-rfc/blob/master/spec.md#3-lnurl-withdraw">выводы через Lightning-ссылки (lnurl)</a> в/из других мест, обрабатывает текущие и ошибочные транзакции, выполняет <a href="https://twitter.com/VNumeris/status/1148403575820709890"> сканирование QR кодов</a> (хотя для этого вы должны сделать фото QR кода своим приложением Telegram и операция может провалиться, в зависимости от модели вашего телефона, терпения и удачи) и другие полезные вещи.

Используя @{{ .BotName }}, вы достаточно хорошо подготовлены для любых взаимодействий с Lightning Network.
    `,
	TUTORIALBLUE: `
Хотя это работает, в реальном мире использование Telegram-чата и вставка счетов может быть проблемным.

Для использования на улицах, вы можете импортировать @{{ .BotName }} средства в <a href="https://bluewallet.io/">BlueWallet</a>. Вам не требуется хранить ваши ончейн Биткоины там, или создавать Lightning кошелёк по умолчанию, вы просто должны ввести команду /bluewallet здесь для импорта URL и вставить её там в окне импорта.

Всё, что вы сделаете в BlueWallet после импорта, будет отражаться в боте, справедливо и обратное (вы будете получать уведомления о платежах в BlueWallet в вашем чате Telegram, но не в кошельке).
    `,
	TUTORIALAPPS: `
Благодаря некоторой магии, которую мы заложили в бот, вы можете без проблем взаимодействовать с внутренними и сторонними приложениями из чата @{{ .BotName }}, автоматически задействовав свой баланс, то есть без опций выбора, ввода количеств (или хуже, счетов Лайтнинг) на вебсайтах, перед тем как провести транзакцию.

Сейчас мы поддерживаем сервисы:

📢 /sats4ads -- получайте оплату за просмотр рекламы, платите за отправку рекламы. /help_sats4ads
☁️ /bitclouds -- создавайте VPS, Bitcoin и Lightning узлы, как сервис. /help_bitclouds
⚽ /microbet -- делайте ставки на microbet.fun и выводите свой баланс в один клик. /help_microbet
🎁 /gifts -- создавайте оплачиваемую ссылку на lightning.gifts, которую вы можете отправить друзьям и получить оповещение, как только купон будет потрачен, не теряйте ссылки востребования средств. /help_gifts
📡 /satellite -- отправляйте сообщения через космический спутник Blockstream. /help_satellite
🎲 /coinflip -- создавайте лотерею типа "победитель получает всё" в групповых чатах. /help_coinflip
🎁 /giveaway  и /giveflip -- создавайте сообщение, которое раздаёт деньги от вас к первому человеку, который нажмёт кнопку или будет выбран случайным образом из нескольких участников. /help_giveaway /help_giveflip
📢 /fundraise -- много людей вкладывают сатоши для одного человека, они наверняка имеют хороший повод. /help_fundraise
📲 /bitrefill -- покупайте подарочные сертификаты и пополняйте телефоны. /help_bitrefill
💸 /yandex and /qiwi -- отправляйте сатоши на счёт yandex.money или qiwi.com в качестве рублей с лучшим курсом.  /help_yandex /help_qiwi 
⛓️ /fundbtc -- отправляйте сатоши с вашего ончейн кошелька Bitcoin на ваш @{{ .BotName }} баланс, сервис обеспечивается golightning.club. /help_fundbtc

Читать больше на странице /help для каждого приложения.
    `,
}
