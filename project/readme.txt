Описание:
Была реализована основная и дополнительная части задачи. Т.к. язык не был
регламентирован, приложение было написано на Golang.
Проблемы:
Постановка задачи была неясна, а именно пункт связанный с нормализацией:
В пункте указано применить нормализацию к фрагментам, полученным в пункте 2
и в качестве примера была предоставлена ссылка на wikipedia, содержащий
алгоритм нормализации серого (greyscaled) изображения. В следствие этого
возникли сомнения в правильности нумерации (основная часть содержит
единственный пункт 1., тогда как дополнительная аналогично начинается с пункта 1. ...,
не должна ли была она начаться с пункта 2?). Поэтому нормализация была
реализована в двух частях: для RGB (Алгоритм взят с недостоверенного сайта) и для Greyscaled
Запуск:
В папке Binaries лежат 4 исполняемых файла, собранных для linux и windows x32 и x64 архитектур.
Работоспособность приложений была проверена в системах:
Ubuntu Budgie x64
Linux Mint x32
Windows 10 x64
Windows 8 x32
Приложение ожидает "увидеть" необходимые папки по пути "./../*", т.е. на уровень ниже.
Однако есть возможность задать флаги с указанием каждой директории при запуске
приложения (Флаги можно посмотреть в исходнике main.go, перечислять не буду, т.к.
вряд ли они понадобятся)
