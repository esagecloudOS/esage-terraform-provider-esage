resource "abiquo_alarm" "test" {
  target      = "${abiquo_virtualmachine.test.id}"
  name        = "test alarm"
  metric      = "cpu_time"
  timerange   = 3
  statistic   = "average"
  formula     = "lessthan"
  threshold   = 10000
}

data "abiquo_virtualdatacenter"        "test"   { name = "tests" }
data "abiquo_template"   "test"   {
  templates = "${data.abiquo_virtualdatacenter.test.templates}"
  name      = "tests"
}

resource "abiquo_virtualappliance" "test" {
  virtualdatacenter = "${data.abiquo_virtualdatacenter.test.id}"
  name              = "test alarm"
}

resource "abiquo_virtualmachine" "test" {
  cpu                    = 1
  deploy                 = false
  ram                    = 64
  label                  = "test alarm"
  virtualappliance       = "${abiquo_virtualappliance.test.id}"
  virtualmachinetemplate = "${data.abiquo_template.test.id}"
}
